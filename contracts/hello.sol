pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

contract Hello {
   string message;

   constructor() public {
      message = "hello from deng mimi";
   }

   event EvtSet(address from, string msg);

   function set(string memory name) public {
      message = name;
      emit EvtSet(msg.sender, message);
   }

   function get() public view returns(string memory, bool) {
      return (message, true);
   }
}