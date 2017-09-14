pragma solidity ^0.4.15;

import "./ERC20Basic.sol";

contract ERC20 is ERC20Basic {
  function allowance(address owner, address spender) constant returns (uint256);
  function transferFrom(address from, address to, uint256 value) returns (bool);
  function approve(address spender, uint256 value) returns (bool);
  function approve(address _owner, address _spender, uint256 _value) returns (bool);
  event Approval(address indexed owner, address indexed spender, uint256 value);
}