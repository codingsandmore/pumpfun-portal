# pumpfun

This is a simple API wrapper for the pump fun portal services. Located here:

https://pumpportal.fun/

## Purpose

to keep track of all newly created pairs and executed prices in a simple, fast fashion.

### Usage

```go
package main

import (
	"github.com/codingsandmore/pumpfun/portal"
	"github.com/codingsandmore/pumpfun/portal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	discoverPair := func(p *portal.NewPairResponse) {
		log.Info().Any("pair", p).Msgf("discovered pair")
	}
	discoverTrade := func(p *portal.NewTradeResponse) {
		log.Info().Any("trade", p).Msg("discovered trade")
	}

	server.NewPortalServer().Discover(discoverPair, discoverTrade)
}

```

This simple example, provides you with a service, which prints out all newly discovered pairs. Once a pair gets created, it will start printing out all trades.

Do with this what you like

example out put
```go
9:13PM INF subscribing to messages decoder={}
9:13PM INF subscribing to messages decoder={}
9:13PM INF attempting to establish connection to wss://pumpportal.fun/api/data
9:13PM INF subscribing to channel for new trades
9:13PM INF connecting to wss://pumpportal.fun/api/data url=wss://pumpportal.fun/api/data
9:13PM INF attempting to establish connection to wss://pumpportal.fun/api/data
9:13PM INF connecting to wss://pumpportal.fun/api/data url=wss://pumpportal.fun/api/data
9:13PM INF sending welcome message welcome="{\"method\" : \"subscribeNewToken\"}"
9:13PM INF discovered pair pair={"bondingCurveKey":"43qMNRPVo1oB8XYc9YZRuN9fKYSn782X5TnreNKWgS5b","initialBuy":59141732.283464,"marketCapSol":31.316014290152186,"mint":"FNzXkQwLWzRm6Z1jJT9gnVH2pHiKiYLw3Ry1btY5KvLU","signature":"4mcbFv4WTEPqbqXF5hN1WksNkzw8H97Y9vhCpv1EeG1Dcj3UMkg26CLNktbQk8J5o53S8ArZGaJTuZukK5aWWQg3","traderPublicKey":"9HkFiXkZgp3fAwNVKDwabLKSvfJExJVDoUpR4Meo7XD7","txType":"create","vSolInBondingCurve":31.749999999999982,"vTokensInBondingCurve":1013858267.716536}
9:13PM INF discovered trade trade={"bondingCurveKey":"43qMNRPVo1oB8XYc9YZRuN9fKYSn782X5TnreNKWgS5b","marketCapSol":32.712100031065454,"mint":"FNzXkQwLWzRm6Z1jJT9gnVH2pHiKiYLw3Ry1btY5KvLU","newTokenBalance":21870594.372929,"signature":"4dyCV1zNQpCPPrKoK5dSH6JEU74j34hrV6VvfhFUVndvZaCJm6nxhRBdPvgVne7PvqTHbquUbvSyMuqDPft1xyhb","tokenAmount":21870594.372929,"traderPublicKey":"7t3Rn6w9ou2kS1cyZjW28G4wkb6N3arLPsqGVogkyQBD","txType":"buy","vSolInBondingCurve":32.44999999999995,"vTokensInBondingCurve":991987673.343607}
9:13PM INF discovered pair pair={"bondingCurveKey":"4A3nTFQ3n81tN3mhSy8uzhZoBzUiEsKi11m8JbnwSqGS","initialBuy":79481481.481481,"marketCapSol":32.6113699906803,"mint":"5XJ77H59Ws9SpEiMj2BjX7fFEcbbGzPowyJBi8ar7i62","signature":"5TKad77QinBX7wpSJABiiH2w7FnSgseEHdjTzTHgT9bcYJAa9xrKyTTfVHifoXXqHZdy1wBfbFHmrVCebBBsfev2","traderPublicKey":"3QH76KPoo47vrwVWecWFYg6q2Pmg2dN4TG3ugDT23UHg","txType":"create","vSolInBondingCurve":32.399999999999984,"vTokensInBondingCurve":993518518.518519}
9:13PM INF discovered pair pair={"bondingCurveKey":"5uVSe536fjF52w3mK5kmVD5rr4kDPAZAdc3aA2fKo7wc","initialBuy":67062499.999999,"marketCapSol":31.811121466293816,"mint":"473DBX5kbuYe8s9Z1ycRJo6y84cCeFEgYp4JcPtJyR4U","signature":"5vDN5Fov9u3ck6YETffjCfoqmK8242cJGMqW8bK6DQb7TYzEtCGG7M3w3Jrq4AJJfxXATZp2VxECHwaY1g1aEbLU","traderPublicKey":"7G6LFz6bNHjMRsvcv6T2sgQ8MHoosJmUUxyjTBF6yk6a","txType":"create","vSolInBondingCurve":31.999999999999968,"vTokensInBondingCurve":1005937500.000001}
9:13PM INF discovered pair pair={"bondingCurveKey":"3xDUW6so9ASWS6e1W2nmD4LheBMX6mSabg6d3dDzimZZ","initialBuy":49470588.235294,"marketCapSol":30.727011494252864,"mint":"CYJz5KPs5drvJVzfregbQ2e8wLUdARq6x6ipPQN1NC9c","signature":"4MEoAhjPa6A6gQVjFvmX1swE4QFM1MS3997gZW276QbzQrmm9HDuBjstJvMBbPfiCJsKS4Vme3b8eKHdTvibtfWZ","traderPublicKey":"GdCV6ncVRcGrXSnuRPUH7dUaaDRWGFfzQPHjEXqutgbk","txType":"create","vSolInBondingCurve":31.449999999999996,"vTokensInBondingCurve":1023529411.764706}
9:13PM INF discovered pair pair={"bondingCurveKey":"BcuB53qRAJJqKQuYVqxKzWSv5SdK37wRDz2NAJCpmNfF","initialBuy":0,"marketCapSol":27.958993476234856,"mint":"FboZUGzat39UewgwBFfQHdBEBiLdKEi2AVnFAbeXrfmh","signature":"JrGoUbLzyYroDwa29qmJ2RyHzUWomtDAwm64yxoCwQzrKftoNx4mBEgFAC5JiYRscye6WaeqBT7j2tagjDsCjpg","traderPublicKey":"Ab9T2tnSrAq6uJ3a5RHx8vFERuQjV55ApNNmzvwScKCr","txType":"create","vSolInBondingCurve":30,"vTokensInBondingCurve":1073000000}
9:13PM INF discovered trade trade={"bondingCurveKey":"43qMNRPVo1oB8XYc9YZRuN9fKYSn782X5TnreNKWgS5b","marketCapSol":29.13456951145911,"mint":"FNzXkQwLWzRm6Z1jJT9gnVH2pHiKiYLw3Ry1btY5KvLU","newTokenBalance":0.283464,"signature":"s29EsThi4if19xueesLZifFhLN1gWYacVqHkGDFLb5yyAVoCSr5iqGvkBUE2t9CiPa9X5uEMnjojstcRVatYo4v","tokenAmount":59141732,"traderPublicKey":"9HkFiXkZgp3fAwNVKDwabLKSvfJExJVDoUpR4Meo7XD7","txType":"sell","vSolInBondingCurve":30.624202725521993,"vTokensInBondingCurve":1051129405.343607}
9:13PM INF discovered pair pair={"bondingCurveKey":"9gZy7eAASMkgTBbP1TnBXP65ysb7CYLyBQznapc886Rn","initialBuy":47840764.33121,"marketCapSol":30.629388008698342,"mint":"Bo51ZyNkhMtQboXq8aUqVPkPeewmCRWR4pF9iq34QSVq","signature":"58de9dg7nQs4cFvRs9cmdDb6i4eR8xtneZNpBRniVufEMEL8iccHeV5aaBdVspr4fB6woRR2Y3z57QU146auoLoY","traderPublicKey":"DniHsMTe5q7SWm6R26bRg6vfrqGvodAxif7JsbmMfyWJ","txType":"create","vSolInBondingCurve":31.399999999999995,"vTokensInBondingCurve":1025159235.66879}

```

Obviously you can work of this example to push the data to a database, a bot or do some magic with it

### Need coffee!!!

we would love to receive your support and would appreciate it, if you send us some SOL for coffee, if you consider this useful :)

```
D79Tjjs6GpsUL2Pxd7PNsWuBaRJs5Qdt1XNorXRD5Azd
```

### Collaboration

Feel free to reach out to me, I'm always interested in some fun collaborations.