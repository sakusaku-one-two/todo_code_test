package todo_usecase

import (
	entity "api/internal/domain/entitys/todo_entity"
	"context"
	"log/slog"
)

type Sort_engin_function_type = func(todos []entity.Todo) ([]entity.Todo, error)

type SortService struct {
	sort_engin Sort_engin_function_type
}

func NewSort(di_sort_engin Sort_engin_function_type) *SortService {
	return &SortService{ //とりあえずヒープに置く
		sort_engin: di_sort_engin,
	}
}

func (ss *SortService) Execute(ctx context.Context, todos []entity.Todo) []entity.Todo {
	result, err := ss.sort_engin(todos)
	if err != nil {
		slog.Log(ctx, slog.LevelError, err.Error()+"at sort function")
		return todos
	}
	return result
}

// 再帰を使用した分割ソート.. さあ　スタックを食い潰そう　(｀・ω・´)
func RecSort(todos []entity.Todo) ([]entity.Todo, error) {

	if is_not_divisible_this(todos) { //分割不可能か判定
		return todos, nil
	}

	median_date := getMedian(todos)          //time.Timeの中央値の取得
	var l_list, r_list, result []entity.Todo // 分割するリスト

	for _, todo := range todos {
		if todo.Limit.GetValue() <= median_date {
			l_list = append(l_list, todo)
		} else {
			r_list = append(r_list, todo)
		}
	}

	l_list, err := RecSort(l_list)
	if err != nil {
		return todos, err
	}
	r_list, err = RecSort(r_list)
	if err != nil {
		return todos, err
	}
	result = append(result, l_list...)
	result = append(result, r_list...)

	return result, nil
}

func is_not_divisible_this(todos []entity.Todo) bool {

}
