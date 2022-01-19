pragma solidity 0.8.11;

contract PoSA {

    constructor () {}

    receive() payable external {}

    struct Signer {
        uint256 stakeAmount;
    }

    mapping(address => Signer) signer;
    uint256 immutable public MINIMUM_STAKE = 100 * 1e18;
    uint256 immutable public SIGNER_AMOUNT = 3;
    address[] public signerList;

    function initialize(address[] memory initialSigners) public {
        updateSigners(initialSigners);
    }

    function getSigners() external view returns (address[] memory) {
        return signerList;
    }

    function updateSigners(address[] memory newSigners) public {
        delete signerList;
        require(newSigners.length <= SIGNER_AMOUNT, "Signers amount exceed limit!");
        for (uint256 i; i < newSigners.length; i++) {
            signerList.push(newSigners[i]);
        }
    }
}