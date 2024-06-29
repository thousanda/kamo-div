package remind

import (
	"testing"
	"time"

	"github.com/thousanda/kamo-div/internal/internal/module/messaging"
	"github.com/thousanda/kamo-div/internal/internal/module/reminder"
)

func TestNewUseCase(t *testing.T) {
	t.Run("意図した通りのRemind use caseが生成される", func(t *testing.T) {
		// 入力を用意
		now := time.Date(2024, 4, 6, 1, 2, 3, 4, time.UTC)
		w := reminder.NewMessageWriter()
		m := messaging.NewDoNothingMessenger()

		// テスト対象の関数を実行
		u := NewUseCase(now, w, m)

		// 結果を検証
		if u.now != now {
			t.Errorf("nowが正しくセットされていません。\n"+
				"got : %v\n"+
				"want: %v", u.now, now)
		}
		if u.reminderMessageWriter != w {
			t.Errorf("reminderMessageWriterが正しくセットされていません。\n"+
				"got : %v\n"+
				"want: %v", u.reminderMessageWriter, w)
		}
		if u.messenger != m {
			t.Errorf("messengerが正しくセットされていません。\n"+
				"got : %v\n"+
				"want: %v", u.messenger, m)
		}
	})
}
