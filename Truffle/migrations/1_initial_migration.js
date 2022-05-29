const Migrations = artifacts.require("Migrations");
var Hello = artifacts.require("./Hello.sol");
var TronToken = artifacts.require("./TronToken.sol");

module.exports = function (deployer) {
  deployer.deploy(TronToken);
};
