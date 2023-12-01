namespace go api

include "dragon.thrift"

service Pandora {
    dragon.DragonSayResponse dragonSay(1: dragon.DragonSayRequest req)
}