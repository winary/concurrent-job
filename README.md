# concurrent-job
A simple scheduler that controls how many jobs can be run simultaneously.

## Example

```go
import "github.com/winary/concurrent-job"

job := cjob.New(2)

for i := 0; i < 5; i++ {
	num := i
	job.Do(func() {
		for i := 0; i < 2; i++ {
			log.Println("==== ", num)
			time.Sleep(time.Second * 3)
		}
	})
}

job.Wait()
```
5 jobs, but only 2 run at the same time.
```
2020/12/12 11:44:13 ====  4
2020/12/12 11:44:13 ====  1
2020/12/12 11:44:16 ====  4
2020/12/12 11:44:16 ====  1
2020/12/12 11:44:19 ====  0
2020/12/12 11:44:19 ====  3
2020/12/12 11:44:22 ====  0
2020/12/12 11:44:22 ====  3
2020/12/12 11:44:25 ====  2
2020/12/12 11:44:28 ====  2
```