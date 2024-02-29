# MODULE Buffer
Buffer module is used on different subproject to normalise and create a secure buffer

This buffer authorize adding item if the size of current folder is less than MaxSize else they stock on folder (pass to NewBuffer) the item for future reading when buffer size will be not fully 
## Installation



```bash
go get github.com/CritsendGo/modBuffer
```


## Usage
```go
import(csBuffer "github.com/CritsendGo/modBuffer")

func main(){
	
    // Create Buffer
    myBuffer:=csBuffer.NewBuffer("/tmp/test-buffer/",64)
	
    // Add Item
    var item time.Time
    item=time.Now()
    err:=myBuffer.Add(item)
    if err!=nil{
        fmt.Println(err)
    }
	
    // Get Next Item
    item,err:=myBuffer.Get()
    if err!=nil{
        fmt.Println(err)
    }
}

```
## License
Attribution-NonCommercial-NoDerivatives 