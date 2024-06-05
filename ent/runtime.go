// Code generated by ent, DO NOT EDIT.

package ent

import (
	"vultisigner/ent/ecdsalocaldata"
	"vultisigner/ent/ecdsapub"
	"vultisigner/ent/pailliersk"
	"vultisigner/ent/schema"
	"vultisigner/ent/vault"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	ecdsapubFields := schema.ECDSAPub{}.Fields()
	_ = ecdsapubFields
	// ecdsapubDescCurve is the schema descriptor for curve field.
	ecdsapubDescCurve := ecdsapubFields[0].Descriptor()
	// ecdsapub.CurveValidator is a validator for the "curve" field. It is called by the builders before save.
	ecdsapub.CurveValidator = func() func(string) error {
		validators := ecdsapubDescCurve.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(curve string) error {
			for _, fn := range fns {
				if err := fn(curve); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	ecdsalocaldataFields := schema.EcdsaLocalData{}.Fields()
	_ = ecdsalocaldataFields
	// ecdsalocaldataDescNTildeI is the schema descriptor for n_tilde_i field.
	ecdsalocaldataDescNTildeI := ecdsalocaldataFields[0].Descriptor()
	// ecdsalocaldata.NTildeIValidator is a validator for the "n_tilde_i" field. It is called by the builders before save.
	ecdsalocaldata.NTildeIValidator = ecdsalocaldataDescNTildeI.Validators[0].(func(string) error)
	// ecdsalocaldataDescH1i is the schema descriptor for h1i field.
	ecdsalocaldataDescH1i := ecdsalocaldataFields[1].Descriptor()
	// ecdsalocaldata.H1iValidator is a validator for the "h1i" field. It is called by the builders before save.
	ecdsalocaldata.H1iValidator = ecdsalocaldataDescH1i.Validators[0].(func(string) error)
	// ecdsalocaldataDescH2i is the schema descriptor for h2i field.
	ecdsalocaldataDescH2i := ecdsalocaldataFields[2].Descriptor()
	// ecdsalocaldata.H2iValidator is a validator for the "h2i" field. It is called by the builders before save.
	ecdsalocaldata.H2iValidator = ecdsalocaldataDescH2i.Validators[0].(func(string) error)
	// ecdsalocaldataDescAlpha is the schema descriptor for alpha field.
	ecdsalocaldataDescAlpha := ecdsalocaldataFields[3].Descriptor()
	// ecdsalocaldata.AlphaValidator is a validator for the "alpha" field. It is called by the builders before save.
	ecdsalocaldata.AlphaValidator = ecdsalocaldataDescAlpha.Validators[0].(func(string) error)
	// ecdsalocaldataDescBeta is the schema descriptor for beta field.
	ecdsalocaldataDescBeta := ecdsalocaldataFields[4].Descriptor()
	// ecdsalocaldata.BetaValidator is a validator for the "beta" field. It is called by the builders before save.
	ecdsalocaldata.BetaValidator = ecdsalocaldataDescBeta.Validators[0].(func(string) error)
	// ecdsalocaldataDescP is the schema descriptor for p field.
	ecdsalocaldataDescP := ecdsalocaldataFields[5].Descriptor()
	// ecdsalocaldata.PValidator is a validator for the "p" field. It is called by the builders before save.
	ecdsalocaldata.PValidator = ecdsalocaldataDescP.Validators[0].(func(string) error)
	// ecdsalocaldataDescQ is the schema descriptor for q field.
	ecdsalocaldataDescQ := ecdsalocaldataFields[6].Descriptor()
	// ecdsalocaldata.QValidator is a validator for the "q" field. It is called by the builders before save.
	ecdsalocaldata.QValidator = ecdsalocaldataDescQ.Validators[0].(func(string) error)
	// ecdsalocaldataDescXi is the schema descriptor for xi field.
	ecdsalocaldataDescXi := ecdsalocaldataFields[7].Descriptor()
	// ecdsalocaldata.XiValidator is a validator for the "xi" field. It is called by the builders before save.
	ecdsalocaldata.XiValidator = ecdsalocaldataDescXi.Validators[0].(func(string) error)
	// ecdsalocaldataDescShareID is the schema descriptor for share_id field.
	ecdsalocaldataDescShareID := ecdsalocaldataFields[8].Descriptor()
	// ecdsalocaldata.ShareIDValidator is a validator for the "share_id" field. It is called by the builders before save.
	ecdsalocaldata.ShareIDValidator = ecdsalocaldataDescShareID.Validators[0].(func(string) error)
	paillierskFields := schema.PaillierSK{}.Fields()
	_ = paillierskFields
	// paillierskDescN is the schema descriptor for n field.
	paillierskDescN := paillierskFields[0].Descriptor()
	// pailliersk.NValidator is a validator for the "n" field. It is called by the builders before save.
	pailliersk.NValidator = paillierskDescN.Validators[0].(func(string) error)
	// paillierskDescLambdaN is the schema descriptor for lambda_n field.
	paillierskDescLambdaN := paillierskFields[1].Descriptor()
	// pailliersk.LambdaNValidator is a validator for the "lambda_n" field. It is called by the builders before save.
	pailliersk.LambdaNValidator = paillierskDescLambdaN.Validators[0].(func(string) error)
	// paillierskDescPhiN is the schema descriptor for phi_n field.
	paillierskDescPhiN := paillierskFields[2].Descriptor()
	// pailliersk.PhiNValidator is a validator for the "phi_n" field. It is called by the builders before save.
	pailliersk.PhiNValidator = paillierskDescPhiN.Validators[0].(func(string) error)
	// paillierskDescP is the schema descriptor for p field.
	paillierskDescP := paillierskFields[3].Descriptor()
	// pailliersk.PValidator is a validator for the "p" field. It is called by the builders before save.
	pailliersk.PValidator = paillierskDescP.Validators[0].(func(string) error)
	// paillierskDescQ is the schema descriptor for q field.
	paillierskDescQ := paillierskFields[4].Descriptor()
	// pailliersk.QValidator is a validator for the "q" field. It is called by the builders before save.
	pailliersk.QValidator = paillierskDescQ.Validators[0].(func(string) error)
	vaultFields := schema.Vault{}.Fields()
	_ = vaultFields
	// vaultDescName is the schema descriptor for name field.
	vaultDescName := vaultFields[0].Descriptor()
	// vault.NameValidator is a validator for the "name" field. It is called by the builders before save.
	vault.NameValidator = vaultDescName.Validators[0].(func(string) error)
	// vaultDescPubKey is the schema descriptor for pub_key field.
	vaultDescPubKey := vaultFields[1].Descriptor()
	// vault.PubKeyValidator is a validator for the "pub_key" field. It is called by the builders before save.
	vault.PubKeyValidator = vaultDescPubKey.Validators[0].(func(string) error)
	// vaultDescLocalPartyKey is the schema descriptor for local_party_key field.
	vaultDescLocalPartyKey := vaultFields[3].Descriptor()
	// vault.LocalPartyKeyValidator is a validator for the "local_party_key" field. It is called by the builders before save.
	vault.LocalPartyKeyValidator = vaultDescLocalPartyKey.Validators[0].(func(string) error)
	// vaultDescChainCodeHex is the schema descriptor for chain_code_hex field.
	vaultDescChainCodeHex := vaultFields[4].Descriptor()
	// vault.ChainCodeHexValidator is a validator for the "chain_code_hex" field. It is called by the builders before save.
	vault.ChainCodeHexValidator = func() func(string) error {
		validators := vaultDescChainCodeHex.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(chain_code_hex string) error {
			for _, fn := range fns {
				if err := fn(chain_code_hex); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
