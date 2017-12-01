package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"xq.goproject.com/commonTools/stringTool"
)

var (
	wg sync.WaitGroup
)

func init() {
	wg.Add(1)
}

func main() {
	playerMgr := newPlayerMgr()

	// insert
	go func() {
		for {
			id := stringTool.GetNewGUID()
			name := fmt.Sprintf("Hero_%s", id)
			obj := newPlayer(id, name)
			playerMgr.insert(obj)

			insert(obj)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// update
	go func() {
		for {
			obj := playerMgr.randomSelect()
			if obj == nil {
				continue
			}
			seed := rand.New(rand.NewSource(time.Now().UnixNano()))
			suffix := seed.Intn(1000)
			newName := fmt.Sprintf("Hero_%d", suffix)
			obj.resetUserName(newName)

			update(obj)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// delete
	go func() {
		for {
			obj := playerMgr.randomSelect()
			if obj == nil {
				continue
			}
			playerMgr.delete(obj)

			clear(obj)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	//	errorFile
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			id := stringTool.GetNewGUID()
			name := fmt.Sprintf("Hero_%s%s", id, id)
			obj := newPlayer(id, name)
			playerMgr.insert(obj)
			print("errorFile")

			insert(obj)
		}

	}()

	wg.Wait()
}
