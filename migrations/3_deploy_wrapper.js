var WrapperOraclize = artifacts.require("./WrapperOraclize.sol");

module.exports = function(deployer) {
 deployer.deploy(WrapperOraclize, {from: web3.eth.accounts[0], gas: 4712380, gasLimit: 200000000});  
};
