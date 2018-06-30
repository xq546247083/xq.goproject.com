package src

import "time"

/*
 * 委托对象
 *
 */

type job func(dtNow time.Time)
