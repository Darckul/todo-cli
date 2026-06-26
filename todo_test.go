package main

import (
	"testing"
)

// TestNextID проверяет, что ID генерируются правильно.
// Это "табличный тест" — популярный Go-паттерн:
// описываем случаи в срезе, гоняем один и тот же код для каждого.
func TestNextID(t *testing.T) {
	cases := []struct {
		name  string
		tasks []Task
		want  int
	}{
		{
			name:  "пустой список — первый ID должен быть 1",
			tasks: []Task{},
			want:  1,
		},
		{
			name:  "одна задача с ID 1 — следующий должен быть 2",
			tasks: []Task{{ID: 1}},
			want:  2,
		},
		{
			name:  "после удаления задачи 2 из [1,2,3] — следующий ID 4, не 3",
			tasks: []Task{{ID: 1}, {ID: 3}},
			want:  4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := nextID(c.tasks)
			if got != c.want {
				t.Errorf("nextID() = %d, хотели %d", got, c.want)
			}
		})
	}
}

// TestSaveAndLoad проверяет, что данные не теряются при записи и чтении.
func TestSaveAndLoad(t *testing.T) {
	// t.TempDir() создаёт временную папку, которая удаляется после теста.
	// Подменяем dataFile, чтобы не трогать реальный tasks.json.
	dataFile = t.TempDir() + "/tasks.json"

	original := []Task{
		{ID: 1, Text: "Тест задача", Done: false},
		{ID: 2, Text: "Выполненная", Done: true},
	}

	if err := save(original); err != nil {
		t.Fatalf("save() вернул ошибку: %v", err)
	}

	loaded, err := load()
	if err != nil {
		t.Fatalf("load() вернул ошибку: %v", err)
	}

	if len(loaded) != len(original) {
		t.Fatalf("load() вернул %d задач, хотели %d", len(loaded), len(original))
	}

	for i, task := range loaded {
		if task.ID != original[i].ID || task.Text != original[i].Text || task.Done != original[i].Done {
			t.Errorf("задача %d не совпадает: got %+v, want %+v", i, task, original[i])
		}
	}
}

// TestLoadEmptyFile проверяет, что при отсутствии файла load() не падает.
func TestLoadEmptyFile(t *testing.T) {
	dataFile = t.TempDir() + "/nonexistent.json"

	tasks, err := load()
	if err != nil {
		t.Fatalf("load() не должен возвращать ошибку для несуществующего файла: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("ожидали пустой список, получили %d задач", len(tasks))
	}
}
