pragma solidity ^0.4.15;

import "./MintableToken.sol";

contract SingleTokenCoin is MintableToken {
    
    string public constant name = "Start mining";
    
    string public constant symbol = "STM";
    
    uint32 public constant decimals = 2;
    
}