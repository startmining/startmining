import React, { Component } from 'react';
import CrowdsaleContract from '../build/contracts/Crowdsale.json';
import getWeb3 from './web3';
import { injectGlobal } from 'styled-components';
const contract = require('truffle-contract');
import 'bootstrap/dist/css/bootstrap.css';
import { Button, Input, Container, Row, Col, Table, Alert, Badge, Card, CardBlock } from 'reactstrap';
import { BootstrapTable, TableHeaderColumn } from 'react-bootstrap-table';
import moment from 'moment';

require("moment-duration-format");

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import getMuiTheme from 'material-ui/styles/getMuiTheme'
import lightBaseTheme from 'material-ui/styles/baseThemes/lightBaseTheme'

const lightMuiTheme = getMuiTheme(lightBaseTheme);

import DatePicker from 'material-ui/DatePicker';
import TimePicker from 'material-ui/TimePicker';

class Crowdsale extends Component {
    constructor(props) {
    super(props);

    this.state = {
      account: 0,
      crowdsaleContract: 0,
      symbol: "STM",

      investorsTokens: 0,
      investorsTokensManual: 0,
      investorsETH: 0,
      priceCurrentInvestor: 0,
      investorsCount: 0,

      minimalPriceCurrentInvestor: 36654609929078010,
      currentTokenPrice: '',

      oraclizeTax: 4000000000000000,

      tokenPricePlaceholder: "Current Token Price",
      buyTokenPlaceholder: "Enter Amount of WEI",

      ether: 1000000000000000000,
      
      eth: 0,
      totalETH: 0,
      balanceContract: 0,

      totalTokens: 0,
      soldTokens: 0,
      leftTokens: 0,

      addressToken: '',
      addressICO: '',

      transactions: [],
      transactionsManual: [],
      transactionsInvestors: [],

      manualInvestorAddress: '',
      manualInvestorToken: '',

      investorAddress: '',
      investorToken: '',

      addressFrom: '',
      addressTo: '',
      addressAmount: 0,

      addressPrivilegedWallet: '',

      dateStart: 0,
      dateFinish: new Date(1508284740 * 1000).toUTCString(),
      dateLeft: 0,

      datePicker: 0,
      startDate: 0,

      addressTransfer: '',
      addressOwner: '',

      alertBuyTokens: false,
      alertBuyTokensFail: false,
    };
  }

  componentWillMount() {
    getWeb3
      .then((results) => {
        this.setState({
          web3: results.web3,
        });
      })
      .then(() => {
        this.setState({account: this.state.web3.eth.accounts[0]})
        
        //Crowdsale
        const crowdsaleContract = contract(CrowdsaleContract);
        crowdsaleContract.setProvider(this.state.web3.currentProvider);

        this.setState({crowdsaleContract})

        this.updateContractInfo();

        this.getCurrentPrice();
        
      })
      .catch((err) => {
        console.log('Error finding web3.', err);
      });
  }

  componentDidMount() {
    setInterval(() => {
     this.updateCountdown();
   }, 1000);
  }

  onDismissBuyAlert() {
    this.setState({ alertBuyTokens: false });
  }

  onShowBuyAlert() {
    this.setState({ alertBuyTokens: true });
  }

  onDismissBuyAlertFail() {
    this.setState({ alertBuyTokensFail: false });
  }

  onShowBuyAlertFail() {
    this.setState({ alertBuyTokensFail: true });
  }

  async getPrice() {
    const { crowdsaleContract, account, oraclizeTax } = this.state;
    this.setState({tokenPricePlaceholder: "Await response...."})    
    const instance = await crowdsaleContract.deployed();
    await instance.sendOracleData({gas: 4412200, from: account, value: oraclizeTax}).then(result => {
      console.log("sendOracleData", result);
      this.setState({tokenPricePlaceholder: "Press Show Price Button"})    
    })

  }

  async showPrice() {
    const { crowdsaleContract, ether, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getWrapperData({gas: 4412200, from: account}).then(result => {
      console.log("showPrice", this.state.web3.toUtf8(result));
      let price = ether / this.state.web3.toUtf8(result);
      this.setState({tokenPricePlaceholder: price / ether + " ETH"})
      
    })

    await instance.getFinishStatus({gas: 4412200, from: account}).then(result => {
      console.log("getFinishStatus", result);
    })


    await instance.getQueryPrice("URL", {gas: 4412200, from: account}).then(result => {
      console.log("getQueryPrice", result);
    })

    await instance.checkWrapperBalance({gas: 4412200, from: account}).then(result => {
      console.log("checkWrapperBalance", result);
    })

    await instance.getBalanceContract({gas: 4412200, from: account }).then(result => {
      console.log("getBalanceContract", result);
    })
  }

  async updateContractInfo() {
    await this.getInvestorInfo();
    await this.getMainInfo();
    await this.getTransactionsByAccount();
    await this.getAddressToken();
    await this.getAddressICO();
    await this.getManualTransactions();
    await this.getICODate();
    await this.getOwner();
    await this.getInvestorsTransactions();
    await this.getInvestorsCount();
  }

  async updateCountdown() {
    
    let now  = moment().format();
    let finish = moment.unix(1508284740).format();
    
    var finishTime = moment(finish);
    var nowTime = moment(now);
    let different = finishTime.diff(nowTime);

    var leftDate = moment.duration(different).format("d [days] h [hours] mm [minutes] ss [seconds]");

    this.setState({dateLeft: leftDate});
  }

  async getTokensByInvestor() {
    const { crowdsaleContract, account } = this.state;

    const instance = await crowdsaleContract.deployed();
    await instance.getInvestorsTokens(account, {gas: 4412200, from: account }).then(result => {
      this.setState({investorsTokens: result.toString()})
    })

    await instance.getTokensManual(account, {gas: 4412200, from: account }).then(result => {
      this.setState({investorsTokensManual: result.toString()})
    })
  }

  async getOwner() {
    const { crowdsaleContract, account } = this.state;

    const instance = await crowdsaleContract.deployed();
    await instance.getOwner({gas: 4412200, from: account }).then(result => {
      this.setState({addressOwner: result.toString()})
    })
  }

  async getAddressToken() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.token({gas: 4412200, from: account }).then(result => {
      this.setState({addressToken: result.toString()})
    })
  }

  async getAddressICO() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getContractAddress({gas: 4412200, from: account }).then(result => {
      console.log("addressICO", result)
      this.setState({addressICO: result.toString()})
    })
  }

  async setTransferOwnership() {
    const { crowdsaleContract, addressTransfer, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.setTransferOwnership(addressTransfer, {gas: 4412200, from: account }).then(result => {
      console.log("setTransferOwnership", result)
      setTimeout(() => {this.updateContractInfo()}, 3000)
    }); 
  }

  async getManualInvestorsCount() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    let count = await instance.getManualInvestorsCount({gas: 4412200, from: account }).then(result => {
      return result.toString();
    });
    return count;
  }

  async getManualAddress(index) {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    let address = await instance.getManualAddress(index, {gas: 4412200, from: account }).then(result => {
      this.setState({manualInvestorAddress: result.toString()})
      return result.toString();
    })
    return address;
  }

  async getTransactionsByAccount (startBlockNumber, endBlockNumber) {
    const { account } = this.state; 
    
    if (endBlockNumber == null) {
        endBlockNumber = await new Promise((resolve, reject) => {
            return this.state.web3.eth.getBlock("latest", function(error, result) {
                if(!error) {
                    resolve(result.number)
                } else {
                    console.error(error)
                }
            })
        }).then(function(data) {
            return data
        })
    }
    if (startBlockNumber == null) {
        startBlockNumber = endBlockNumber - 100;
    }
    let array = [];
    for (let i = startBlockNumber; i <= endBlockNumber; i++) {
        var block = await new Promise((resolve, reject) => {
            return this.state.web3.eth.getBlock(i, true, function(error, result) {
                if(!error) {
                    resolve(result)
                } else {
                    console.error(error)
                }
            })
        }).then(function(data) {
            return data
        })
        if (block != null && block.transactions != null) {
            block.transactions.forEach(async (e) => {
                if (account == e.from || account == e.to) {
                    array.push(e);
                }
            })
        }
    }
    this.setState({transactions: array});
  }

  async getManualTransactions() {
    
    await this.getManualInvestorsCount().then(async (dataCount) => {
      let manualArray = [];    
      for (let i = 0; i < dataCount; i++) {
        
        await this.getManualAddress(i).then(async (dataAddress) => {
          await this.getManualByAddress(dataAddress).then((dataTokens) => {
            manualArray.push({
              to: dataAddress,
              value: dataTokens / 100
            });
          });
        });
      }
      this.setState({transactionsManual: manualArray});
      
    });
    
  }

  async getInvestorsTransactions() {
    const { crowdsaleContract, account } = this.state; 
        let instance = await crowdsaleContract.deployed();
        await instance.getInvestors({gas: 2000000, from: account }).then(async (dataCount) => {

          let countInv = dataCount.toString();
          const investors = [];
    
          for (let i = 0; i < countInv; i++) {
            await instance.getInvestorByIndex(i, {gas: 4712200, from: account }).then(async (dataAddress) => {
            console.log("INVESTORS_ADDRESS",dataAddress.toString());            
              let address = dataAddress.toString();

              await instance.getInvestorByValue(address, {gas: 4712200, from: account }).then(async (dataETH) => {
                console.log("INVESTORS_VALUE",dataETH.toString());
                investors.push({
                  address: address,
                  count: dataETH.toString() 
                });
              })
            })
          };
          this.setState({ transactionsInvestors: investors });
        });

        console.log("-------//////-----", this.state.transactionsInvestors)
      }

  async getInvestorsCount() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getInvestors({gas: 4412200, from: account }).then((result) => {
      this.setState({investorsCount: result.toString()})
    })
  }

  async getCurrentPrice() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getCurrentPrice({gas: 4412200, from: account }).then((result) => {
      this.setState({priceCurrentInvestor: result.toString()})
    })
  }

  async finishMinting() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.finishMinting({gas: 4412200, from: account }).then(result => {
      console.log("finishMinting", result)
      setTimeout(() => {this.updateContractInfo()}, 3000)
    });
  }

  async buyTokens() {
    const { crowdsaleContract, account, eth } = this.state;
    this.setState({buyTokenPlaceholder: "Await response...."})    
    const instance = await crowdsaleContract.deployed();
    await instance.mint({gas: 4412200, from: account, value: eth })
    .then((result) => {
      console.log("buyTokens", result)
      setTimeout(() => {this.getInvestorInfo(); this.getInvestorsCount(); this.getInvestorsTransactions(); this.getMainInfo()}, 3000);
      this.setState({buyTokenPlaceholder: "Congratulations!"})
    });
  }

  async manualRefund() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.manualRefund({gas: 4412200, from: account })
    .then((result) => {
      console.log("manualRefund", result)
      setTimeout(() => {this.getInvestorInfo(); this.getInvestorsCount(); this.getInvestorsTransactions(); this.getMainInfo()}, 3000);
    });
  }

  async getSoldToken() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getSoldToken({gas: 4412200, from: account }).then(result => {
      this.setState({soldTokens: result.toString()});
    })
  }

  async getLeftToken() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getLeftToken({gas: 4412200, from: account }).then(result => {
      this.setState({leftTokens: result.toString()})
    })
  }

  async getTotalToken() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getTotalToken({gas: 4412200, from: account }).then(result => {
      this.setState({totalTokens: result.toString()})
    })
  }

  async getTotalETH() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getTotalETH({gas: 4412200, from: account }).then(result => {
      this.setState({totalETH: result.toString()})
    })
  }

  async getBalanceContract() {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    await instance.getBalanceContract({gas: 4412200, from: account }).then(result => {
      this.setState({balanceContract: result.toString()})
    })
  }

  async sendToAddress() {
    const { crowdsaleContract, manualInvestorAddress, manualInvestorToken, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    
    await instance.sendToAddress(manualInvestorAddress, manualInvestorToken * 100, {gas: 4412200, from: account }).then(result => {
      console.log("sendToAddress", result)
      setTimeout(() => {this.getInvestorInfo(); this.getManualTransactions(); this.getMainInfo()}, 3000)
    });
  }

  async getManualByAddress(manualAddress) {
    const { crowdsaleContract, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    
    let tokens = await instance.getManualByAddress(manualAddress, {gas: 4412200, from: account }).then(result => {
      return result.toString();
    })
    return tokens;
  }

  async transferTokens() {
    const { crowdsaleContract, addressFrom, addressTo, addressAmount, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    
    await instance.transfer(addressFrom, addressTo, addressAmount, {gas: 4412200, from: account }).then(result => {
      console.log("transfer", result)
    })
  }

  async setPrivlegedWallet() {
    const { crowdsaleContract, addressPrivilegedWallet, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    console.log("addressPrivilegedWallet", addressPrivilegedWallet);
    await instance.setPrivelegedWallet(addressPrivilegedWallet, {gas: 4412200, from: account }).then(result => {
      console.log("setPrivlegedWallet", result)
    })
  }

  async removePrivlegedWallet() {
    const { crowdsaleContract, addressPrivilegedWallet, account } = this.state;
    
    const instance = await crowdsaleContract.deployed();
    
    await instance.removePrivelegedWallet(addressPrivilegedWallet, {gas: 4412200, from: account }).then(result => {
      console.log("removePrivlegedWallet", result)
    })
  }

  async setStartDate() {
    const { crowdsaleContract, account, startDate } = this.state; 
    const instance = await crowdsaleContract.deployed();

    let date = startDate / 1000;
    console.log("--------", date)

    await instance.setICODate(date, {gas: 2012200, from: account }).then(result => {
      console.log("setICODate", result)
      setTimeout(() => {this.updateContractInfo()}, 3000)
      });
  }

  async getICODate() {
    const { crowdsaleContract, account } = this.state; 
    const instance = await crowdsaleContract.deployed();

    await instance.getICODate({gas: 2012200, from: account }).then(result => {
      this.setState({dateStart: new Date(result.toString() * 1000).toUTCString()})
      
      })
  }

  getTransactionByAddress() {
    return(
      <BootstrapTable height='330' scrollTop={ 'Bottom' } data={this.state.transactions} striped hover condensed>
        <TableHeaderColumn width='20%' dataAlign="center" dataField="hash" dataFormat={ txFormatter } isKey >HASH</TableHeaderColumn>
        <TableHeaderColumn width='20%' dataAlign="center" dataField="from" >FROM</TableHeaderColumn>
        <TableHeaderColumn width='20%' dataAlign="center" dataField="to" >TO</TableHeaderColumn>
        <TableHeaderColumn width='20%' dataAlign="center" dataField="value" >VALUE</TableHeaderColumn>
    </BootstrapTable>
    );
  }

  getTransactionByManualSend() {
    return(
      <BootstrapTable height='270' scrollTop={ 'Bottom' } data={this.state.transactionsManual} striped hover condensed>
        <TableHeaderColumn width='20%' dataField="to" isKey >To</TableHeaderColumn>
        <TableHeaderColumn width='20%' dataField="value" >Value</TableHeaderColumn>
    </BootstrapTable>
    );
  }

  getTransactionByInvestor() {
    return(
      <BootstrapTable height='270' scrollTop={ 'Bottom' } data={this.state.transactionsInvestors} striped hover condensed>
        <TableHeaderColumn width='20%' dataFormat={ linkFormatter } dataField="address" isKey >From</TableHeaderColumn>
        <TableHeaderColumn width='20%' dataField="count" >Value</TableHeaderColumn>
    </BootstrapTable>
    );
  }

  async getInvestorInfo() {
    await this.getTokensByInvestor();
    await this.getCurrentPrice();
  }

  async getMainInfo() {
    await this.getSoldToken();
    await this.getLeftToken();
    await this.getTotalETH();
    await this.getTotalToken();
  }

  showTimePicker(x, event) {
    this.refs.timepicker.openDialog();
    this.setState({datePicker: event})
  }

  showDatePicker() {
    this.refs.datepicker.openDialog();
  }

  setTimePicker(time) {
    let hour = time.getHours();
    let minute = time.getMinutes();
    let timePicker = moment.utc(this.state.datePicker).add({hours: hour, minutes: minute})
    this.setState({startDate: timePicker._d})
  }

    render() {
      return (
        <Container>
          <br></br>
          <br></br>
          <h1>Ethereum STM Token ICO</h1>
          <Col>
            <Row>
              <Col md={{ size: '5'}}>
                <br></br>
                <br></br> 
                <br></br> 
                <h3>Contract Info</h3>
                <hr color="blue" className="my-3" />                
                <br></br> 
                <Table striped bordered>
                  <thead>
                    <tr>
                      <th>Info</th>
                      <th>Value</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>Address ICO</td>
                      <td>{this.state.addressICO}</td>
                    </tr>
                    <tr>
                      <td>Address Token</td>
                      <td>{this.state.addressToken}</td>
                    </tr>
                    <tr>
                      <td>Address Owner</td>
                      <td>{this.state.addressOwner}</td>
                    </tr>
                    <tr>
                      <td>Date Start</td>
                      <td>{this.state.dateStart}</td>
                    </tr>
                    <tr>
                      <td>Date Finish</td>
                      <td>{this.state.dateFinish}</td>
                    </tr>
                    <tr>
                      <td>Date Left</td>
                      <td>{this.state.dateLeft}</td>
                    </tr>
                  </tbody>
                </Table>
                <br></br> 
                <br></br> 
                <h3>Investor Info</h3>
                <hr color="blue" className="my-3" />                                
                <br></br> 
                <Table striped bordered>
                  <thead>
                    <tr>
                      <th>Info</th>
                      <th>Value</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>Investor Token (For ETH)</td>
                      <td>{this.state.investorsTokens / 100} {this.state.symbol}</td>
                    </tr>
                    <tr>
                      <td>Investor Token (Manual)</td>
                      <td>{this.state.investorsTokensManual / 100} {this.state.symbol}</td>
                    </tr>
                    <tr>
                      <td>Price for minimum amount (10 {this.state.symbol})</td>
                      <td>{this.state.minimalPriceCurrentInvestor} WEI</td>
                    </tr>
                  </tbody>
                </Table>
                <br></br> 
                <br></br> 
                <h3>Main Info</h3>
                <hr color="blue" className="my-3" />                
                <br></br> 
                <Table striped bordered>
                  <thead>
                    <tr>
                      <th>Info</th>
                      <th>Value</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>Total Tokens</td>
                      <td>{this.state.totalTokens / 100}</td>
                    </tr>
                    <tr>
                      <td>Sold Tokens</td>
                      <td>{this.state.soldTokens / 100}</td>
                    </tr>
                    <tr>
                      <td>Left Tokens</td>
                      <td>{this.state.leftTokens / 100}</td>
                    </tr>
                    <tr>
                      <td>Total ETH</td>
                      <td>{this.state.totalETH}</td>
                    </tr>
                    <tr>
                      <td>Investors Count</td>
                      <td>{this.state.investorsCount}</td>
                    </tr>
                  </tbody>
                </Table>
                <br></br> 
                <br></br>                 
                <br></br>
                <br></br>
                <h3>Manual Send</h3>
                <hr color="blue" className="my-3" />                
                <br></br> 
                {this.getTransactionByManualSend()}
              </Col>
              <Col md={{ size: '5', offset: 2 }}>
                <br></br>
                <br></br> 
                <Row>
                  <Col>
                    <Col md={{ size: '3', offset: 5 }}>
                      <h6><Badge color="success" pill>for production</Badge></h6>                      
                    </Col>
                    <h3>User Functions</h3>
                    <hr color="success" className="my-3" />
                  </Col>
                </Row>
                <br></br>
                <div>
                  <Card style={{ backgroundColor: 'whitesmoke' }}>
                    <CardBlock>
                      <h5>Get Current Token Price</h5>                
                      <Input 
                        value={this.state.currentTokenPrice}
                        placeholder={this.state.tokenPricePlaceholder}
                      />
                      <br></br>
                      <Row>
                        <Col md={{ size: '3' }}>
                          <Button color="info" onClick={() => this.getPrice()} >Get Price</Button>
                        </Col>
                        <Col md={{ size: '3', offset: 1 }}>
                          <Button color="warning" onClick={() => this.showPrice()} >Show Price</Button>
                        </Col>
                      </Row>
                    </CardBlock>
                </Card>
                </div>
                <br></br>   
                <div>
                  <Card style={{ backgroundColor: 'whitesmoke' }}>
                    <CardBlock>
                      <h5>Enter Amount of WEI</h5>                
                      <Input 
                        value={this.state.eth}
                        placeholder={this.state.buyTokenPlaceholder}
                        onChange={e => this.setState({ eth: e.target.value })}
                        onKeyDown={this.handleSubmit}
                      />
                      <br></br>
                      <Row>
                        <Col md={{ size: '3' }}>
                          <Button color="success" onClick={() => this.buyTokens()} >Buy Tokens</Button>
                        </Col>
                        <Col md={{ size: '3', offset: 1 }}>
                          <Button color="danger" onClick={() => this.manualRefund()} >Refund Tokens</Button>
                        </Col>
                      </Row>
                    </CardBlock>
                </Card>
                </div>
                <br></br>              
                <br></br> 
                <Row>
                  <Col>
                    <Col md={{ size: '3', offset: 5 }}>
                      <h6><Badge color="success" pill>for production</Badge></h6>                      
                    </Col>
                    <h3>Admin Functions</h3>
                    <hr color="success" className="my-3" />
                  </Col>
                </Row>
                <Row>
                <Col md={{ size: '12' }}>
                  <br></br>
                  <div>
                    <Card style={{ backgroundColor: 'whitesmoke' }}>
                      <CardBlock>
                        <h4>Granted Wallets</h4>                
                        <br></br>              
                        <h5>Enter Address to privelege Wallet</h5>                                  
                        <Input 
                          value={this.state.addressPrivilegedWallet}
                          placeholder="Enter Address to privilege Wallet"
                          onChange={e => this.setState({ addressPrivilegedWallet: e.target.value })}
                          onKeyDown={this.handleSubmit}
                        />
                        <br></br>
                        <Row>
                          <Col md={{ size: '3' }}>
                            <Button color="success" onClick={() => this.setPrivlegedWallet()} >Grant Right</Button>
                          </Col>
                          <Col md={{ size: '3', offset: 1 }}>
                            <Button color="danger" onClick={() => this.removePrivlegedWallet()} >Revork Right</Button>
                          </Col>
                        </Row>
                      </CardBlock>
                    </Card>
                  </div>
                  <br></br>                
                  </Col>
                </Row>
                <br></br> 
                <Row>
                <Col md={{ size: '12' }}>
                  <div>
                    <Card style={{ backgroundColor: 'whitesmoke' }}>
                      <CardBlock>
                        <h4>Manual send Token to Investor</h4>                
                        <br></br>
                        <h5>Enter Investor's Address</h5>                
                        <Input 
                          value={this.state.manualInvestorAddress}
                          placeholder="Enter Investor's Address"
                          onChange={e => this.setState({ manualInvestorAddress: e.target.value })}
                          onKeyDown={this.handleSubmit}
                        />
                        <br></br>
                        <h5>Enter Amount of Token for Investor</h5>                                  
                        <Input 
                          value={this.state.manualInvestorToken}
                          placeholder="Enter Amount of Token for Investor"
                          onChange={e => this.setState({ manualInvestorToken: e.target.value})}
                          onKeyDown={this.handleSubmit}
                        />
                        <br></br>
                        <Row>
                          <Col md={{ size: '3' }}>
                            <Button color="primary" onClick={() => this.sendToAddress()} >Send</Button>
                          </Col>
                        </Row>
                      </CardBlock>
                    </Card>
                  </div>
                  <br></br>
                  </Col>
                </Row>
                <br></br>
                <Row>
                <Col md={{ size: '12' }}>
                  <div>
                    <Card style={{ backgroundColor: 'whitesmoke' }}>
                      <CardBlock>
                        <h4>Transfer Ownership</h4>                
                        <br></br>
                        <h5>Enter New Owner Address</h5>                
                        <Input 
                          value={this.state.addressTransfer}
                          placeholder="Enter New Owner Address"
                          onChange={e => this.setState({ addressTransfer: e.target.value })}
                          onKeyDown={this.handleSubmit}
                        />
                        <br></br>
                        <Row>
                          <Col md={{ size: '3' }}>
                            <Button color="warning" onClick={() => this.setTransferOwnership()} >Send</Button>
                          </Col>
                        </Row>
                      </CardBlock>
                    </Card>
                  </div>
                  <br></br>
                  <br></br>
                  </Col>
                </Row>
                <h3>Investors</h3>
                <hr color="blue" className="my-3" />                
                <br></br> 
                {this.getTransactionByInvestor()}
              </Col>
              </Row>
              <Row>
                <Col>
                  <br></br>
                  <br></br>
                  <Row>
                    <Col>
                      <Col md={{ size: '3', offset: 2 }}>
                        <h6><Badge color="danger" pill>only for Demo version</Badge></h6>                      
                      </Col>
                      <h3>Test & Debug</h3>
                      <hr color="red" className="my-3" />
                    </Col>
                  </Row>
                  <br></br>
                  <Row>
                    <Col md={{ size: '5' }}>
                      <div>
                          <Card style={{ backgroundColor: 'whitesmoke' }}>
                            <CardBlock>
                              <h5>Enter Date for Start ICO</h5>
                                <Input 
                                  value={this.state.startDate}
                                  placeholder="Date for start ICO"
                                  onFocus={() => this.showDatePicker()}
                                  onKeyDown={this.handleSubmit}
                                />
                                <MuiThemeProvider muiTheme={lightMuiTheme}>
                                  <div>
                                    <DatePicker
                                      id="dataPickerId"
                                      onChange={(x, event) => {this.showTimePicker(x, event)}  }
                                      ref="datepicker"
                                      style={{ display: 'none' }}
                                    />
                                    <TimePicker
                                      id="timePickerId"
                                      ref="timepicker"
                                      style={{ display: 'none' }}
                                      onChange={(x, event) => {this.setTimePicker(event)}}                    
                                    />
                                  </div>
                                </MuiThemeProvider>
                                <br></br>
                                <Row>              
                                  <Col md={{ size: '3'}}>
                                    <Button color="success" onClick={() => this.setStartDate()} >
                                      Start ICO
                                    </Button>
                                  </Col>
                                </Row>
                            </CardBlock>
                          </Card>
                        </div>
                        <br></br>
                        <Row>
                          <Col md={{ size: '3' }}>
                            <Button color="danger" onClick={() => this.finishMinting()} >Finish Minting</Button>
                          </Col>
                        </Row>
                      </Col>
                    <br></br>
                    <br></br>
                      <Col md={{ size: '5', offset: 2 }}>
                        <div>
                          <Card style={{ backgroundColor: 'whitesmoke' }}>
                            <CardBlock>
                              <h4>Transfer Tokens</h4>
                              <br></br>                  
                              <Row>    
                                <Col md={{ size: '4' }}>   
                                  <h5>From</h5>                                               
                                  <Input 
                                    value={this.state.addressFrom}
                                    placeholder="Enter address of Investor (From)"
                                    onChange={e => this.setState({ addressFrom: e.target.value })}
                                    onKeyDown={this.handleSubmit}
                                  />
                                </Col>
                                <Col md={{ size: '4' }}>                           
                                  <h5>To</h5>                                               
                                  <Input 
                                    value={this.state.addressTo}
                                    placeholder="Enter address of Investor (To)"
                                    onChange={e => this.setState({ addressTo: e.target.value })}
                                    onKeyDown={this.handleSubmit}
                                  />
                                </Col>
                                <Col md={{ size: '4' }}>                           
                                  <h5>Amount</h5>                                               
                                  <Input 
                                    value={this.state.addressAmount}
                                    placeholder="Enter Amount of Tokens"
                                    onChange={e => this.setState({ addressAmount: e.target.value })}
                                    onKeyDown={this.handleSubmit}
                                  />
                                </Col>
                              </Row>
                              <br></br>
                              <Row>
                                <Col md={{ size: '4' }}>
                                  <Button color="warning" onClick={() => this.transferTokens()} >Transfer</Button>
                                </Col>
                              </Row>
                          </CardBlock>
                        </Card>
                      </div>
                    </Col>
                  </Row>
                  <br></br>
                  <br></br>
                </Col>
              </Row>
              <br></br> 
              <br></br> 
              <h3>Transactions</h3>  
              {this.getTransactionByAddress()}
          </Col>
        </Container>
      );
    }
}

export default Crowdsale;

function txFormatter(cell, row) {
  return '<tr><td><a href="https://etherscan.io/tx/' + cell + '" target="_blank">' + cell + '</a></td><tr>';
}

function linkFormatter(cell, row) {
  return '<tr><td><a href="https://etherscan.io/address/' + cell + '" target="_blank">' + cell + '</a></td><tr>';
}

injectGlobal`
  @import url('https://fonts.googleapis.com/css?family=Roboto');

  body {
    background-color: whitesmoke;
    font-family: 'Roboto', sans-serif;
  }
`
