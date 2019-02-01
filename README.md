# 如何使用cosmos-sdk建立一个Nameservice

## 背景介绍
* `Nameservice`，即域名注册系统是一个比较流行的区块链应用领域。开发人员已经尝试在不同的区块链平台上实现其功能，代表产品有： Namecoin, ENS, 和Handshake。人们可以借助这些产品实现传统的DNS域名的买卖。
* Tendermint
Tendermint是一个复制状态机，用来实现区块链应用的网络层和共识层的功能。简而言之，Tendermint负责处理应用所处理的交易的顺序并在全网广播。Tendermint也是拜占庭容错(BFT)的。
* Cosmos-SDK
Cosmos-SDK是区块链应用开发框架。它具有模块化的特性，每个模块都包含自己的消息/事务处理器。各类消息经过指定的路由进入到对应的模块。

## 设计NameService
以下是`nameservice`应用程序所需的模块:
* `auth`: 该模块用于定义帐户和费用，并确定是交易数据否有访问权限。
* `bank`: 该模块用于建立并管理代币账户。
* `nameservice`: 该模块就是需要我们去构建的！

## 主要设计目标：状态和消息类型
1. 状态：状态实时反映区块链应用的运行情况。我们可以从状态中了解到各个账户的余额还有域名的所有者和报价，以及域名对应的IP。auth 和 bank 是Cosmos-SDK的基础模块。它们的功能已经被实现。所以我们主要目的在于实现NameService的内容。
在Cosmos-SDK中，所有数据都被存储在multistore中。multistore由一些KVStores组成。表示了键值对 key/value。对于一个NameService应用你需要设计以下键值对：
通过nameStore存储 name 到value的映射 
通过ownerStore存储 name 到owner的映射
通过priceStore存储 name 到price的映射

2. 消息：消息包含在交易中。它们会触发应用状态变化。每个模块都定义了一个消息列表以及其Handler。 以下是为nameservice应用实现其功能所需的消息：
MsgSetName: 该消息允许域名所有者在nameStore中为其设定一个IP。
MsgBuyName: 该消息允许客户购买一个域名并在ownerStore中成为其所有者。

当交易（包含在块中）到达Tendermint节点时，它将通过ABCI 传递给应用程序并解码以获取消息。 然后将消息发送到适当的模块，并根据Handler中定义的逻辑进行处理。 如果需要更新状态，则Handler会调用Keeper来执行更新。 本教程的后续步骤可了解更多相关概念的信息。