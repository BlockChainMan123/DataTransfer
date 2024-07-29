contract LoadPhase {

    //defined the following variables that are used to transfer pd
    uint public TransferDatapdIndex;
    mapping(address => mapping(uint => plaintextData)) transferDatapdStores;
    mapping(uint => address) transferdatapdIdSrote;

    //defined the following variables that are used to transfer cd
    uint public TransferDatacdIndex;
    mapping(address => mapping(uint => ciphertextData)) transferDatacdStores;
    mapping(uint => address) transferdatacdIdSrote;

    //defined the following variables that are used to transfer signature
    uint public TransfersignatureIndex;
    mapping(address => mapping(uint => signature)) transfersignatureStores;
    mapping(uint => address) transfersignatureIdSrote;
    
    //defined the following variables that are used to transfer t
    uint public rIndex;
    mapping(address => mapping(uint => rofsignature)) rStores;
    mapping(uint => address) rIdSrote;

    //defined the following variables that are used to load ct
    uint public ctIndex;
    mapping(address => mapping(uint => loadct)) ctStores;
    mapping(uint => address) ctIdSrote;

    //defined the following variables that are used to load ct and the 1st-signature
    uint public ctFirstSignatureIndex;
    mapping(address => mapping(uint => loadctFirstSignature)) ctFirstSignatureStores;
    mapping(uint => address) ctFirstSignatureIdSrote;

    //defined the following variables that are used to load 2cd-signature
    uint public SecondSignatureIndex;
    mapping(address => mapping(uint => loadSecondSignature)) SecondSignatureStores;
    mapping(uint => address) SecondSignatureIdSrote;

    

    //a plaintext data struct
    struct plaintextData{//all fields
        uint  id;          // plaintext data ID
        string signature;  // signatures
        string pk;        // address asociated to pk
        string data;        // the transferred data
    }

    //a ciphertext data struct
    struct ciphertextData{//all fields
        uint  id;          // plaintext data ID
        string signature;  // signatures
        string pk;        // address asociated to pk
        string firstsignature;      // the first signature
        string cd;        // the transferred data in the form of ciphertext data has been saved on source chain.
    }

    //signature data struct
    struct signature{//all fields
        uint  id;          // plaintext data ID
        string signature;  // signatures
    }

    //deliver r to smart contrct
    struct rofsignature{//all fields
        uint  id;          // ct ID
        string r;          // r of the first signature
    }

    //load ct to smart contrct
    struct loadct{//all fields
        uint  id;          // ct ID
        string ct;          // ct associated to pd
    }

    //load ct and 1st-signature to smart contrct
    struct loadctFirstSignature{//all fields
        uint  id;          // ct ID
        string ct;          // ct associated to pd
        string firstsignature;          // the first signature
    }

    //load 2st-signature to smart contrct
    struct loadSecondSignature{//all fields
        uint  id;          // ct ID
        string secondsignature;          // the second signature
    }

    //hash the ct & its id
    function getMessageHash(bytes32 _ct,uint _accountId) public pure returns(bytes32){
        return keccak256(abi.encodePacked(_ct,_accountId));
    }

    //Sign on a hash to ct
    function toEthSignedMessageHash(bytes32 hash) public pure returns (bytes32) {
        // 哈希的长度为32
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
     
    //Deliver pd along with its proof 
    function deliverDatapdProof(string memory _signature, string memory _pk, string memory _data) public {
        TransferDatapdIndex += 1;
        plaintextData memory PlaintextData = plaintextData(TransferDatapdIndex, _signature, _pk, _data);
        transferDatapdStores[msg.sender][TransferDatapdIndex] = PlaintextData;
        transferdatapdIdSrote[TransferDatapdIndex] = msg.sender;
    }

    //committee downloads the DatapdProof
    function getDatapdProof(uint id) view public returns(uint, string memory, string memory, string memory){
       plaintextData memory PlaintextData = transferDatapdStores[transferdatapdIdSrote[id]][id];
       return(PlaintextData.id, PlaintextData.signature, PlaintextData.pk, PlaintextData.data);
   }  

    //Deliver cd along with its proof
    function deliverDatacdProof(string memory _signature, string memory _pk, string memory _firstsignature, string memory _cd) public {
        TransferDatacdIndex += 1;
        ciphertextData memory CiphertextData = ciphertextData(TransferDatacdIndex, _signature, _pk, _firstsignature, _cd);
        transferDatacdStores[msg.sender][TransferDatacdIndex] = CiphertextData;
        transferdatacdIdSrote[TransferDatacdIndex] = msg.sender;
    }

    //committee downloads the DatacdProof
    function getDatacdProof(uint id) view public returns(uint, string memory, string memory, string memory, string memory){
       ciphertextData memory CiphertextData = transferDatacdStores[transferdatacdIdSrote[id]][id];
       return(CiphertextData.id, CiphertextData.signature, CiphertextData.pk, CiphertextData.firstsignature, CiphertextData.cd);
   }

   //committee loads ct
    function loadctdata(string memory _ct) public {
        ctIndex += 1;
        loadct memory Loadct = loadct(ctIndex, _ct);
        ctStores[msg.sender][ctIndex] = Loadct;
        ctIdSrote[ctIndex] = msg.sender;
    }

    //committee loads ct and 1st-signature
    function loadctFirstSignaturedata(string memory _ct, string memory _firstsignature) public {
        ctFirstSignatureIndex += 1;
        loadctFirstSignature memory LoadctFirstSignature = loadctFirstSignature(ctFirstSignatureIndex, _ct, _firstsignature);
        ctFirstSignatureStores[msg.sender][ctFirstSignatureIndex] = LoadctFirstSignature;
        ctFirstSignatureIdSrote[ctFirstSignatureIndex] = msg.sender;
    }

    //committee loads 2cd-signature
    function loadSecondSignaturedata(string memory _secondsignature) public {
        SecondSignatureIndex += 1;
        loadSecondSignature memory LoadSecondSignature = loadSecondSignature(SecondSignatureIndex, _secondsignature);
        SecondSignatureStores[msg.sender][SecondSignatureIndex] = LoadSecondSignature;
        SecondSignatureIdSrote[SecondSignatureIndex] = msg.sender;
    }

    //doctor retrieves ct
    function retrievect(uint id) view public returns(uint, string memory){
       loadct memory Loadct = ctStores[ctIdSrote[id]][id];
       return(Loadct.id, Loadct.ct);
    }

    //doctor retrieves ct and the first signature
    function retrievectandsignature(uint id) view public returns(uint, string memory, string memory){
       loadctFirstSignature memory LoadctFirstSignature = ctFirstSignatureStores[ctFirstSignatureIdSrote[id]][id];
       return(LoadctFirstSignature.id, LoadctFirstSignature.ct, LoadctFirstSignature.firstsignature);
   }

    //doctor retrieves 2cd-signature
    function retrieve2signature(uint id) view public returns(uint, string memory){
       loadSecondSignature memory LoadSecondSignature = SecondSignatureStores[SecondSignatureIdSrote[id]][id];
       return(LoadSecondSignature.id, LoadSecondSignature.secondsignature);
    }

   //doctor delivers r of the first signature to smart contract
    function deliverrofsignature(string memory _signature) public {
        rIndex += 1;
        rofsignature memory Rofsignature = rofsignature(rIndex, _signature);
        rStores[msg.sender][rIndex] = Rofsignature;
        rIdSrote[rIndex] = msg.sender;
    }

    //Deliver signature the second signature in the second round.
    function deliversignature(string memory _signature) public {
        TransfersignatureIndex += 1;
        signature memory Signature = signature(TransfersignatureIndex, _signature);
        transfersignatureStores[msg.sender][TransfersignatureIndex] = Signature;
        transfersignatureIdSrote[TransfersignatureIndex] = msg.sender;
    }

    //committee retrieves the second signature
    function retrievesecondsignature(uint id) view public returns(uint, string memory){
       signature memory Signature = transfersignatureStores[transfersignatureIdSrote[id]][id];
       return(Signature.id, Signature.signature);
   }

   

}
