// SPDX-License-Identifier: MIT
pragma solidity ^0.8.14;

contract RobotSwarm {
    uint256 public total;

    function getOrder() public returns (string memory) {
        if (total > 10) {
            total = 0;
            return "Right";
        }
        return "Left";
    }

    function increment(uint256 number) public {
        total += number;
    }
}
