# Data

## Data目录和Biz目录的关系
data目录采用了依赖倒置原则，在biz目录定义接口，data目录实现接口，并依赖接口。
由data目录去实现接口。这样的好处是，biz目录不需要关心data目录的具体实现，只需要关心接口即可。

## Data目录怎么实现的
通过ent orm框架