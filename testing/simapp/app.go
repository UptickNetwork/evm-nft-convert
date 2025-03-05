package app

import (
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/runtime"
	srvflags "github.com/evmos/ethermint/server/flags"
	"github.com/evmos/ethermint/x/evm"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/evmos/ethermint/x/evm/vm/geth"
	"github.com/evmos/ethermint/x/feemarket"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"cosmossdk.io/simapp"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/evidence"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/nft"
	abci "github.com/cometbft/cometbft/abci/types"
	tmos "github.com/cometbft/cometbft/libs/os"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/ibc-go/modules/capability"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	//nftkeeper "cosmossdk.io/x/nft/keeper"
	//nftmodule "cosmossdk.io/x/nft/module"
	simappparams "cosmossdk.io/simapp/params"
	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	solomachine "github.com/cosmos/ibc-go/v8/modules/light-clients/06-solomachine"
	"github.com/cosmos/ibc-go/v8/testing/simapp/upgrades"
	ethermint "github.com/evmos/ethermint/types"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"

	"github.com/UptickNetwork/evm-nft-convert"
	erc721keeper "github.com/UptickNetwork/evm-nft-convert/keeper"
	erc721types "github.com/UptickNetwork/evm-nft-convert/types"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	"github.com/UptickNetwork/evm-nft-convert/testing/simapp/ante"
)

const appName = "SimApp"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	feemarketParams = feemarkettypes.Params{
		NoBaseFee:                false,
		BaseFeeChangeDenominator: 8,
		ElasticityMultiplier:     4,
		BaseFee:                  math.NewInt(10000000000),
		MinGasPrice:              math.LegacyNewDecFromInt(math.NewInt(10000000000)),
		MinGasMultiplier:         math.LegacyNewDecWithPrec(5, 1),
	}

	legacyProposalHandlers = []govclient.ProposalHandler{
		paramsclient.ProposalHandler,
	}

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			legacyProposalHandlers,
		),
		groupmodule.AppModuleBasic{},
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		solomachine.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		vesting.AppModuleBasic{},
		consensus.AppModuleBasic{},
		evm.AppModuleBasic{},
		erc721.AppModuleBasic{},
		feemarket.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		nft.ModuleName:                 nil,

		evmtypes.ModuleName:    {authtypes.Minter, authtypes.Burner}, // used for secure addition and subtraction of balance using module account
		erc721types.ModuleName: nil,
	}
)

var (
	//	_ simapp.SimApp           = (*SimApp)(nil)
	_ servertypes.Application = (*SimApp)(nil)
)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimApp struct {
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	txConfig          client.TxConfig

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    *stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	GroupKeeper      groupkeeper.Keeper
	CrisisKeeper     *crisiskeeper.Keeper
	UpgradeKeeper    *upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	AuthzKeeper      authzkeeper.Keeper

	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedFeeMockKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedIBCMockKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAMockKeeper       capabilitykeeper.ScopedKeeper
	ScopedNFTTransferKeeper   capabilitykeeper.ScopedKeeper

	// Ethermint keepers
	EvmKeeper       *evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper
	// Erc721 keepers
	Erc721Keeper erc721keeper.Keeper

	// the module manager
	mm *module.Manager
	bm module.BasicManager
	// simulation manager
	sm *module.SimulationManager

	// the configurator
	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".simapp")
	// manually update the power reduction by replacing micro (u) -> atto (a) uptick
	sdk.DefaultPowerReduction = ethermint.PowerReduction
}

// NewSimApp returns a reference to an initialized SimApp.
func NewSimApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig simappparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SimApp {
	appCodec := encodingConfig.Codec
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(
		appName,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		crisistypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		group.StoreKey,
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		capabilitytypes.StoreKey,
		authzkeeper.StoreKey,
		consensusparamtypes.StoreKey,

		// ethermint keys
		evmtypes.StoreKey,
		feemarkettypes.StoreKey,
		erc721types.StoreKey,
	)
	tkeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey, evmtypes.TransientKey, feemarkettypes.TransientKey)
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &SimApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
		txConfig:          encodingConfig.TxConfig,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// get authority address
	//authAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	//// set the BaseApp's parameter store
	//app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, keys[consensusparamtypes.StoreKey], authAddr)
	//bApp.SetParamStore(&app.ConsensusParamsKeeper)

	// set the BaseApp's parameter store
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[consensusparamtypes.StoreKey]),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		runtime.EventService{},
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(&app.ConsensusParamsKeeper.ParamsStore)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	// seal capability keeper after scoping modules
	app.CapabilityKeeper.Seal()

	// SDK module keepers

	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[authtypes.StoreKey]),
		ethermint.ProtoAccount,
		maccPerms,
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[banktypes.StoreKey]),
		app.AccountKeeper,
		BlockedAddresses(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		logger,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[stakingtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)

	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[minttypes.StoreKey]),
		app.StakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[distrtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		legacyAmino,
		runtime.NewKVStoreService(app.keys[slashingtypes.StoreKey]),
		app.StakingKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	// Create Ethermint keepers
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec,
		authtypes.NewModuleAddress(govtypes.ModuleName),
		keys[feemarkettypes.StoreKey],
		tkeys[feemarkettypes.TransientKey],
		app.GetSubspace(feemarkettypes.ModuleName),
	)

	tracer := cast.ToString(appOpts.Get(srvflags.EVMTracer))
	app.EvmKeeper = evmkeeper.NewKeeper(
		appCodec,
		keys[evmtypes.StoreKey],
		tkeys[evmtypes.TransientKey],
		authtypes.NewModuleAddress(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.FeeMarketKeeper,
		nil,
		geth.NewEVM,
		tracer,
		app.GetSubspace(evmtypes.ModuleName),
	)

	app.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[crisistypes.StoreKey]),
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[feegrant.StoreKey]),
		app.AccountKeeper,
	)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(app.keys[upgradetypes.StoreKey]),
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(
		runtime.NewKVStoreService(app.keys[authzkeeper.StoreKey]),
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	// IBC Keepers
	// register the proposal types
	govRouter := govv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper))

	govConfig := govtypes.DefaultConfig()
	/*
		Example of setting gov params:
		govConfig.MaxMetadataLen = 10000
	*/

	govKeeper := govkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[govtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		bApp.MsgServiceRouter(),
		govConfig,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// Set legacy router for backwards compatibility with gov v1beta1
	govKeeper.SetLegacyRouter(govRouter)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)

	groupConfig := group.DefaultConfig()
	/*
		Example of setting group params:
		groupConfig.MaxMetadataLen = 1000
	*/
	app.GroupKeeper = groupkeeper.NewKeeper(
		keys[group.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
		groupConfig,
	)

	// Create Interchain Accounts Stack
	// SendPacket, since it is originating from the application to core IBC:
	// icaAuthModuleKeeper.SendTx -> icaController.SendPacket -> fee.SendPacket -> channel.SendPacket

	// initialize ICA module with mock module as the authentication module on the controller side

	// create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(app.keys[evidencetypes.StoreKey]),
		app.StakingKeeper,
		app.SlashingKeeper,
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.ProvideCometInfoService(),
	)

	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		// SDK app modules
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp,
			encodingConfig.TxConfig,
		),

		auth.NewAppModule(
			appCodec,
			app.AccountKeeper,
			authsims.RandomGenesisAccounts,
			app.GetSubspace(authtypes.ModuleName),
		),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(
			appCodec,
			app.BankKeeper,
			app.AccountKeeper,
			app.GetSubspace(banktypes.ModuleName),
		),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		crisis.NewAppModule(
			app.CrisisKeeper,
			skipGenesisInvariants,
			app.GetSubspace(crisistypes.ModuleName),
		),
		feegrantmodule.NewAppModule(
			appCodec,
			app.AccountKeeper,
			app.BankKeeper,
			app.FeeGrantKeeper,
			app.interfaceRegistry,
		),
		gov.NewAppModule(
			appCodec,
			&app.GovKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.GetSubspace(govtypes.ModuleName),
		),
		mint.NewAppModule(
			appCodec,
			app.MintKeeper,
			app.AccountKeeper,
			nil,
			app.GetSubspace(minttypes.ModuleName),
		),
		slashing.NewAppModule(
			appCodec,
			app.SlashingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
			app.GetSubspace(slashingtypes.ModuleName),
			app.interfaceRegistry,
		),
		distr.NewAppModule(
			appCodec,
			app.DistrKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
			app.GetSubspace(distrtypes.ModuleName),
		),
		staking.NewAppModule(
			appCodec,
			app.StakingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.GetSubspace(stakingtypes.ModuleName),
		),
		upgrade.NewAppModule(app.UpgradeKeeper, authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		authzmodule.NewAppModule(
			appCodec,
			app.AuthzKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		groupmodule.NewAppModule(
			appCodec,
			app.GroupKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),
		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),
		// Ethermint app modules
		evm.NewAppModule(app.EvmKeeper, app.AccountKeeper, app.GetSubspace(evmtypes.ModuleName)),
		feemarket.NewAppModule(app.FeeMarketKeeper, app.GetSubspace(feemarkettypes.ModuleName)),
		// Uptick app modules
		erc721.NewAppModule(app.Erc721Keeper, app.AccountKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		group.ModuleName,
		consensusparamtypes.ModuleName,
		nft.ModuleName,

		feemarkettypes.ModuleName,
		evmtypes.ModuleName,
		erc721types.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		group.ModuleName,
		consensusparamtypes.ModuleName,
		nft.ModuleName,

		feemarkettypes.ModuleName,
		evmtypes.ModuleName,
		erc721types.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		group.ModuleName,
		consensusparamtypes.ModuleName,
		nft.ModuleName,

		feemarkettypes.ModuleName,
		evmtypes.ModuleName,
		erc721types.ModuleName,
	)

	app.mm.RegisterInvariants(app.CrisisKeeper)
	app.configurator = module.NewConfigurator(
		app.appCodec,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
	)
	app.mm.RegisterServices(app.configurator)

	// add test gRPC service for testing gRPC queries in isolation
	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(
			appCodec,
			app.AccountKeeper,
			authsims.RandomGenesisAccounts,
			app.GetSubspace(authtypes.ModuleName),
		),
		bank.NewAppModule(
			appCodec,
			app.BankKeeper,
			app.AccountKeeper,
			app.GetSubspace(banktypes.ModuleName),
		),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		feegrantmodule.NewAppModule(
			appCodec,
			app.AccountKeeper,
			app.BankKeeper,
			app.FeeGrantKeeper,
			app.interfaceRegistry,
		),
		gov.NewAppModule(
			appCodec,
			&app.GovKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.GetSubspace(govtypes.ModuleName),
		),
		mint.NewAppModule(
			appCodec,
			app.MintKeeper,
			app.AccountKeeper,
			nil,
			app.GetSubspace(minttypes.ModuleName),
		),
		staking.NewAppModule(
			appCodec,
			app.StakingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.GetSubspace(stakingtypes.ModuleName),
		),
		distr.NewAppModule(
			appCodec,
			app.DistrKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
			app.GetSubspace(distrtypes.ModuleName),
		),
		slashing.NewAppModule(
			appCodec,
			app.SlashingKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
			app.GetSubspace(slashingtypes.ModuleName),
			app.interfaceRegistry,
		),

		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		authzmodule.NewAppModule(
			appCodec,
			app.AuthzKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.interfaceRegistry,
		),

		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		evm.NewAppModule(app.EvmKeeper, app.AccountKeeper, app.GetSubspace(evmtypes.ModuleName)),
		feemarket.NewAppModule(app.FeeMarketKeeper, app.GetSubspace(feemarkettypes.ModuleName)),
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	maxGasWanted := cast.ToUint64(appOpts.Get(srvflags.EVMMaxTxGasWanted))
	options := ante.HandlerOptions{
		AccountKeeper:   app.AccountKeeper,
		BankKeeper:      app.BankKeeper,
		FeeMarketKeeper: app.FeeMarketKeeper,
		EvmKeeper:       app.EvmKeeper,
		FeegrantKeeper:  app.FeeGrantKeeper,
		SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
		SigGasConsumer:  SigVerificationGasConsumer,
		MaxTxGasWanted:  maxGasWanted,
	}

	if err := options.Validate(); err != nil {
		panic(err)
	}

	// app.SetAnteHandler(anteHandler)
	app.SetAnteHandler(ante.NewAnteHandler(options))
	app.SetEndBlocker(app.EndBlocker)
	app.setupUpgradeHandlers()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// Name returns the name of the App
func (app *SimApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *SimApp) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	return app.mm.BeginBlock(ctx)
}

// EndBlocker application updates every end block
func (app *SimApp) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return app.mm.EndBlock(ctx)
}

// InitChainer application update at chain initialization
func (app *SimApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	var genesisState simapp.GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	fmt.Printf("xxl genesisState %v \n ", genesisState)
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *SimApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *SimApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// GetModuleManager returns the app module manager
// NOTE: used for testing purposes
func (app *SimApp) GetModuleManager() *module.Manager {
	return app.mm
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns SimApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns SimApp's InterfaceRegistry
func (app *SimApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *SimApp) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// TestingApp functions

// GetBaseApp implements the TestingApp interface.
func (app *SimApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *SimApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *SimApp) GetTxConfig() client.TxConfig {
	return app.txConfig
}

// SimulationManager implements the SimulationApp interface
func (app *SimApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	cmtservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *SimApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(
		app.BaseApp.GRPCQueryRouter(),
		clientCtx,
		app.BaseApp.Simulate,
		app.interfaceRegistry,
	)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *SimApp) RegisterTendermintService(clientCtx client.Context) {
	cmtservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

func (app *SimApp) RegisterNodeService(clientCtx client.Context, c config.Config) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter(), c)
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(ctx client.Context, rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// ModuleAccountAddrsLegacy returns all the app's module account addresses.
func ModuleAccountAddrsLegacy() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range GetMaccPerms() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func BlockedAddresses() map[string]bool {
	modAccAddrs := ModuleAccountAddrsLegacy()

	// allow the following addresses to receive funds
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(
	appCodec codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	key, tkey storetypes.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	// SDK subspaces
	// paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(authtypes.ModuleName).WithKeyTable(authtypes.ParamKeyTable())
	//paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName).WithKeyTable(banktypes.ParamKeyTable())
	// paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName).WithKeyTable(stakingtypes.ParamKeyTable())
	// paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName).WithKeyTable(minttypes.ParamKeyTable())
	// paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName).WithKeyTable(distrtypes.ParamKeyTable())
	// paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName).WithKeyTable(slashingtypes.ParamKeyTable())
	//paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	// paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName).WithKeyTable(crisistypes.ParamKeyTable())
	// ethermint subspaces
	// paramsKeeper.Subspace(evmtypes.ModuleName)
	paramsKeeper.Subspace(evmtypes.ModuleName).WithKeyTable(evmtypes.ParamKeyTable())
	// paramsKeeper.Subspace(feemarkettypes.ModuleName).WithKeyTable(feemarkettypes.ParamKeyTable())
	paramsKeeper.Subspace(feemarkettypes.ModuleName).WithKeyTable(feemarkettypes.ParamKeyTable())
	// uptick subspaces

	paramsKeeper.Subspace(erc721types.ModuleName).WithKeyTable(erc721types.ParamKeyTable())

	return paramsKeeper
}

// setupUpgradeHandlers sets all necessary upgrade handlers for testing purposes
func (app *SimApp) setupUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(
		upgrades.V5,
		upgrades.CreateDefaultUpgradeHandler(app.mm, app.configurator),
	)

}
