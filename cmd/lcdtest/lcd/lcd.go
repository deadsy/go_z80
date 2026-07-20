//-----------------------------------------------------------------------------
/*

Character LCD Routines

*/
//-----------------------------------------------------------------------------

package lcd

import "errors"

//-----------------------------------------------------------------------------

type HD44780 interface {
	ReadCommand() byte
	WriteCommand(cmd byte)
	ReadData() byte
	WriteData(val byte)
}

type LCD struct {
	dev        HD44780
	rows, cols int
}

func New(dev HD44780, rows, cols int) (*LCD, error) {
	return &LCD{
		dev:  dev,
		rows: rows,
		cols: cols,
	}, nil
}

func (lcd *LCD) checkRowCol(row, col int) error {
	if row < 0 || row >= lcd.rows {
		return errors.New("bad row")
	}
	if col < 0 || col >= lcd.cols {
		return errors.New("bad column")
	}
	return nil
}

//-----------------------------------------------------------------------------

func (lcd *LCD) SetCursor(row, col int) error {
	err := lcd.checkRowCol(row, col)
	if err != nil {
		return err
	}

	return nil
}

func (lcd *LCD) String(row, col int, s string) error {
	err := lcd.checkRowCol(row, col)
	if err != nil {
		return err
	}
	err = lcd.SetCursor(row, col)
	if err != nil {
		return err
	}
	n := min(len(s), lcd.cols-col)
	for i := 0; i < n; i++ {
		lcd.dev.WriteData(s[i])
	}
	return nil
}

//-----------------------------------------------------------------------------
