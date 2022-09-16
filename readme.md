# Taylor 서버
Taylor 서버 Repository

# 블록체인을 이용한 티켓팅 경매 어플, Taylor 🏘️

<p align="center">  

## 👩🏻‍💻👨🏻‍💻 Server 팀원 소개

<div align="center">
  <table>
    <tr>
      <td>
        <a href="https://github.com/xhaktmchl">
          <img src="https://avatars.githubusercontent.com/u/62229967?v=4" width=200/>
          <br>
          <center>Developer-토마스최/최태규</center>
        </a>
      </td>
      <td>
        <a href="https://github.com/Jodongjin">
          <img src="https://user-images.githubusercontent.com/87434861/176502256-cedcd53c-c895-427f-8579-fd53e3ce4f20.jpg" width=200/>
          <br>
          <center>Developer-애플/조동진</center>
        </a>
      </td>
      <td>
        <a href="https://github.com/hyeong-jun-kim">
          <img src= "https://user-images.githubusercontent.com/84059402/186575691-b4dee57b-92f4-4193-9b58-67382939419c.png" width=200/>
          <br>
          <center>Developer-네오/김형준</center>
        </a>
      </td>
    </tr>
  </table>
</div>

## 핵심 기능 ⚙️


### 티켓 경매 기능

티켓 판매자가 티켓을 등록하면 구매자가 원하는 가격을 제시한다. 제시 된 가격이 채택되면 블록체인에 티켓 가격과 소유자를 기록한다.


## 1차 목표 기능 🔥

- **티켓 블록체인 생성 및 티켓 목록 조회**
- **경매 시 채택되면 블록체인에 티켓 가격과 소유자를 업데이트하는 기능**
  <br>

## Goal 🎯

긱사생 어플의 `주요 타겟층`은 대학교 기숙사에서 거주하고 있는 기숙사 생입니다.  
기숙사에서 거주하고있는 유저의 특징은 각 사용자간 거리가 밀접하고, 같은 학교이기 때문에 일반 거주자보다 신뢰성이 높다는 특징을 가지고 있습니다.  
이런 특징들은, 배달 음식 구인 기능, 심부름 및 거래기능에 대한 접근성의 제한이 낮지 않다는 것을 의미합니다.  
위 근거로 앞서 말한 기능을 구현할 예정이며 디벨롭 목표로 `런칭 목표`는 **학교 기숙사 생끼리 배달 음식을 공유할 수 있는 체계**를 구현하는 것이고,  
사용자 반응을 보면서 심부름 기능, 거래기능 및 커뮤니티 기능을 디벨롭 하여 기숙사생이면 필수로 설치해야 하는 어플로 나아가게 할 계획입니다!

테일러 어플의 `주요 타겟층`은 선착순으로 티켓팅하는 대신에 금액을 더 지불하더라도 티켓을 구하고자 하는 사용자들 입니다.
테일러 어플의 기대효과는 티켓을 판매하는 티켓을 판매하는 주최측은 더 높은 수익을 올릴 수 있고, 구매자는 원하는 티켓을 구매해서 서로의 수요를 충족시키는 것입니다.
부가적인 효과는 선착순이라는 점을 이용해 티켓을 대거 구입해서 불법적으로 가격을 높여 이득을 취하는 현상을 일부 억제할 수 있습니다. 

# =================================================================
# paper
**paper** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/paper@latest! | sudo bash
```
`username/paper` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
