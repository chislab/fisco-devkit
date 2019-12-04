pragma solidity ^0.5.13;

contract HelloWorld {
    string name;

    constructor() public {
       name = "Hi, fisco devkit!";
    }

    function get() public view returns(string memory) {
        return name;
    }

    function set(string memory n) public {
    	name = n;
    }
}
