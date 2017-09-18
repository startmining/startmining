pragma solidity ^0.4.15;

import "./StandardToken.sol";
import "./Ownable.sol";
import "./SafeMath.sol";

contract MintableToken is StandardToken, Ownable {

  using SafeMath for uint256;

  bool mintingFinished = false;

  uint256 mintedTokens = 0;

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

  modifier canMint() {
    require(!mintingFinished);
    _;
  }

  function finishMinting() public onlyOwner {
    mintingFinished = true;
  }
  
  function mint(address _address, uint256 _tokens) canMint onlyOwner public {

    require(mintedTokens < totalSupply);

    Mint(_address, _tokens);

    balances[_address] = balances[_address].add(_tokens);

    mintedTokens = mintedTokens.add(_tokens);
  }

  function burnTokens(address _address) onlyOwner public {
    balances[_address] = 0;
    totalSupply = 0;
    mintedTokens = 0;
  }

  function burnFinish() onlyOwner public {
    totalSupply = mintedTokens;
  }

}