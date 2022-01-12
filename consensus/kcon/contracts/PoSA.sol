pragma solidity 0.8.11;

contract PoSA {

    constructor () {
        MINIMUM_STAKE = 100 * 1e18;
        SIGNER_AMOUNT = 3;
    }

    receive() payable external {}

    struct Signer {
        uint256 stakeAmount;
    }

    mapping(address => Signer) signer;
    uint256 immutable public MINIMUM_STAKE;
    uint256 immutable public SIGNER_AMOUNT;
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