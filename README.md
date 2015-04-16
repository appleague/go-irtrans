# go-irtrans
go script for irtrans device

#help
list remotes and cmds: go run go-irtrans.go list
run cmd for remote: go run go-irtrans.go remote cmd

#examples
#get list of remotes and cmds
go run go-irtrans.go list
denon:off,on,tv,airplay,xbox360,atv2,volup,voldown,mute,source,sound_standard,sound_direct,sound_simulation,menu,chlevel,search,return

#send 'volup' ir-code to avr
go run go-irtrans.go denon volup
