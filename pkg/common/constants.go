package common

const DefaultPhone = "18199996701"
const DefaultPassword = "1234"
const NamePrefix = "user_"

const UserApiMachineId = 1
const ContentRpcMachineId = 2
const Twepoch = int64(1483228800000)                  //开始时间截 (2017-01-01)
const WorkeridBits = uint(10)                         //机器id所占的位数
const SequenceBits = uint(12)                         //序列所占的位数
const WorkeridMax = int64(-1 ^ (-1 << WorkeridBits))  //支持的最大机器id数量
const SequenceMask = int64(-1 ^ (-1 << SequenceBits)) //
const WorkeridShift = SequenceBits                    //机器id左移位数
const TimestampShift = SequenceBits + WorkeridBits    //时间戳左移位数
