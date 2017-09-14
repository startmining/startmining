pragma solidity ^0.4.15;

import "./StandardToken.sol";
import "./Ownable.sol";
import "./SafeMath.sol";

contract MintableToken is StandardToken, Ownable {

  using SafeMath for uint256;

  event Mint(address indexed to, uint256 amount);

  event MintFinished();

  event ShowInfo(uint256 _info, string _message);

  function setTotalSupply(uint256 _amount) public onlyOwner returns(uint256) {
    totalSupply = _amount;
    return totalSupply;
  }

  function getTotalTokenCount() public constant returns(uint256) {
    return totalSupply;
  }
  
  function mint(address _address, uint256 _tokens) public {

    Mint(_address, _tokens);

    balances[_address] = balances[_address].add(_tokens);
  }

  function burnTokens(address _address) public {
    balances[_address] = 0;
    totalSupply = 0;
  }

}