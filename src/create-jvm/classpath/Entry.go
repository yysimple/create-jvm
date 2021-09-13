package classpath

import "os"

// :(linux/unix) or ;(windows)这里是用来读取系统的分割符的
const pathListSeparator = string(os.PathListSeparator)

// Entry /*  实现类路径：可以把类路径想象成一个大的整体，它由启动类路径、扩展类路径和用户类路径三个小路径构成。三个小路径又分别由更小的路径构成*/
type Entry interface {
}
