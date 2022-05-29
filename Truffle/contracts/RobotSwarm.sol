// SPDX-License-Identifier: MIT
pragma solidity ^0.8.14;

contract RobotSwarm {
    uint256 public total;

    function getOrder() public view returns (string memory) {
        if (total % 2 == 0) {
            return "Right";
        }
        return "Left";
    }

    function increment(uint256 number) public {
        total += number;
    }
}
