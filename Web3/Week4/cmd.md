1. 什么是merge block
    ```text
    Merge Block 指的是以太坊完成“合并（The Merge）”事件时的那个关键区块。
    
    “The Merge” 是指以太坊从 PoW（工作量证明） 向 PoS（权益证明） 的共识机制转换的事件。
    
    Merge Block 就是 以太坊主网（执行层） 与 信标链（共识层） 合并时的那个区块。
    
    📌 特点：
    
    Merge Block 是区块链历史中唯一的、不重复的区块。
    
    它是 PoW 链的最后一个区块，也是 PoS 链的起点。
    
    Merge 发生在 2022 年 9 月，以太坊从那时起不再依赖挖矿。
    ```

2. 什么是checkpoint
   ```text
   在区块链中，Checkpoint 是一种防止回滚（重组）攻击、提高同步效率的机制，指的是某些被“确认”的历史区块，被标记为不可更改，起到一个“锚点”作用。
   
   用途：
   安全性：防止攻击者通过重组历史区块来篡改数据。
   
   同步优化：新节点同步区块时可以从 checkpoint 开始，而不必从创世块全量同步。
   
   例子（以太坊）：
   信标链每 32 个 epoch 会生成一个finalized checkpoint，一旦 finality 达成，该 checkpoint 之前的链数据不可更改。
   
   轻节点同步（light client）时只需要获取最近的 checkpoint，而不需要下载整个链。
   ```
   

