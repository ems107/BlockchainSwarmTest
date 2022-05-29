const Migrations = artifacts.require("Migrations");
var RobotSwarm = artifacts.require("./RobotSwarm.sol")

module.exports = function (deployer) {
  deployer.deploy(RobotSwarm);
};
