// Code generated by "stringer -type TShirtSize"; DO NOT EDIT.

package main

import "fmt"

const _TShirtSize_name = "BoySBoyMBoyLBoyXLBoyXXLGirlSGirlMGirlLTankSTankMTankLTankXL"

var _TShirtSize_index = [...]uint8{0, 4, 8, 12, 17, 23, 28, 33, 38, 43, 48, 53, 59}

func (i TShirtSize) String() string {
	i -= 1
	if i < 0 || i >= TShirtSize(len(_TShirtSize_index)-1) {
		return fmt.Sprintf("TShirtSize(%d)", i+1)
	}
	return _TShirtSize_name[_TShirtSize_index[i]:_TShirtSize_index[i+1]]
}

const _BoolOrEmpty_name = "OuiNon"

var _BoolOrEmpty_index = [...]uint8{0, 3, 6}

func (i BoolOrEmpty) String() string {
	i -= 1
	if i >= BoolOrEmpty(len(_BoolOrEmpty_index)-1) {
		return fmt.Sprintf("BoolOrEmpty(%d)", i+1)
	}
	return _BoolOrEmpty_name[_BoolOrEmpty_index[i]:_BoolOrEmpty_index[i+1]]
}

const _JobsType_name = "AcceuilArtisteAcceuilPublicBacklineCaisseEcocupEnvironmentRestaurationMerchandisingMontageRunsBarman"

var _JobsType_index = [...]uint8{0, 14, 27, 35, 41, 47, 58, 70, 83, 90, 94, 100}

func (i JobsType) String() string {
	i -= 1
	if i < 0 || i >= JobsType(len(_JobsType_index)-1) {
		return fmt.Sprintf("JobsType(%d)", i+1)
	}
	return _JobsType_name[_JobsType_index[i]:_JobsType_index[i+1]]
}

const _EnglishLevel_name = "NonUn peuScolaireBonFluentBilingue"

var _EnglishLevel_index = [...]uint8{0, 3, 9, 17, 20, 26, 34}

func (i EnglishLevel) String() string {
	i -= 1
	if i < 0 || i >= EnglishLevel(len(_EnglishLevel_index)-1) {
		return fmt.Sprintf("EnglishLevel(%d)", i+1)
	}
	return _EnglishLevel_name[_EnglishLevel_index[i]:_EnglishLevel_index[i+1]]
}

const _EmergencyContactType_name = "DadMumFamillyPartnerOther"

var _EmergencyContactType_index = [...]uint8{0, 3, 6, 13, 20, 25}

func (i EmergencyContactType) String() string {
	i -= 1
	if i < 0 || i >= EmergencyContactType(len(_EmergencyContactType_index)-1) {
		return fmt.Sprintf("EmergencyContactType(%d)", i+1)
	}
	return _EmergencyContactType_name[_EmergencyContactType_index[i]:_EmergencyContactType_index[i+1]]
}

const _Regime_name = "ToutVegetarienVegan"

var _Regime_index = [...]uint8{0, 4, 14, 19}

func (i Regime) String() string {
	i -= 1
	if i < 0 || i >= Regime(len(_Regime_index)-1) {
		return fmt.Sprintf("Regime(%d)", i+1)
	}
	return _Regime_name[_Regime_index[i]:_Regime_index[i+1]]
}
