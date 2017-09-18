pragma solidity ^0.4.15;

contract Addresses {

    //2%
    address public bounty;
    
    //5%
    address public successFee;

    //93%
    address public addr1;

    //93%
    address public addr2;

    //93%
    address public addr3;

    //93%
    address public addr4;


  function Addresses() {

      //2%       //ORIGINAL
      bounty = 0x0064952457905eBFB9c0292200A74B1d7414F081;
                 //TEST
   //   bounty = 0x1626079328312cdb1e731a934a547c6d81b3ee2c;
      
      //5%       //ORIGINAL
      successFee = 0xdA39e0Ce2adf93129D04F53176c7Bfaaae8B051a;
                 //TEST
    //  successFee = 0xf280dacf47f33f442cf5fa9d20abaef4b6e9ca12;

    //93%       //ORIGINAL
      addr1 = 0x300b848558DC06E32658fFB8D59C859D0812CA6C;

      //93%       //ORIGINAL
      addr2 = 0x4388AD192b0DaaDBBaa86Be0AE7499b8D44C5f75;

      //93%       //ORIGINAL
      addr3 = 0x40C9E2D0807289b4c24B0e2c34277BDd7FaCfd87;

      //93%       //ORIGINAL
      addr4 = 0x4E3B219684b9570D0d81Cc13E5c0aAcafe2323B1;
      

     /* //93%       //TEST
      addr1 = 0x1626079328312cdb1e731a934a547c6d81b3ee2c;

      //93%       //TEST
      addr2 = 0x1626079328312cdb1e731a934a547c6d81b3ee2c;

      //93%       //TEST
      addr3 = 0x1626079328312cdb1e731a934a547c6d81b3ee2c;

      //93%       //TEST
      addr4 = 0x1626079328312cdb1e731a934a547c6d81b3ee2c;*/
  }

}
