# MODULE Buffer
Buffer module is used on different subproject to normalise and create a secure buffer

## Installation



```bash
go get github.com/CritsendGo/modBuffer
```


## Usage
```go
import(csBuffer "github.com/CritsendGo/modBuffer")

func main(){
    myBuffer:=csBuffer.NewBuffer("/tmp/test-buffer/",64)
}

```
## License
Attribution-NonCommercial-NoDerivatives 