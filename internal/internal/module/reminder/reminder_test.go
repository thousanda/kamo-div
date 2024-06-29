package reminder

import "testing"

func TestNewReminder(t *testing.T) {
	t.Run("Reminderを作成できる", func(t *testing.T) {
		// テスト対象のものを用意
		s := []string{"a", "b", "c"}

		// テスト対象のメソッドを実行
		r := NewReminder(s)

		// 結果を検証
		if len(r) != len(s) {
			t.Errorf("lengthが違います。\n"+
				"got : %v\n"+
				"want: %v", len(r), len(s))
		}
		for i := range r {
			if r[i] != s[i] {
				t.Errorf("%d番目の要素が違います。\n"+
					"got : %v\n"+
					"want: %v", i, r[i], s[i])
			}
		}
	})
}

func TestReminder_ToStrings(t *testing.T) {
	t.Run("Reminderをstringのスライスにできる", func(t *testing.T) {
		// テスト対象のものを用意
		r := Reminder{"a", "b", "c"}

		// テスト対象のメソッドを実行
		got := r.ToStrings()

		// 結果を検証
		want := []string{"a", "b", "c"}
		if len(got) != len(want) {
			t.Errorf("lengthが違います。\n"+
				"got : %v\n"+
				"want: %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("%d番目の要素が違います。\n"+
					"got : %v\n"+
					"want: %v", i, got[i], want[i])
			}
		}
	})
}
