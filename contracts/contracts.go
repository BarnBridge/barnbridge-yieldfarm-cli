//go:generate abigen --abi ../abi/Staking.json --pkg contracts --out Staking.go --type Staking
//go:generate abigen --abi ../abi/YieldFarm.json --pkg contracts --out YieldFarm.go --type YieldFarm
//go:generate abigen --abi ../abi/YieldFarmLP.json --pkg contracts --out YieldFarmLP.go --type YieldFarmLP
package contracts
