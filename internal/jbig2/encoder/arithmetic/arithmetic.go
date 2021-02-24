//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package arithmetic ;import (_b "bytes";_a "github.com/unidoc/unipdf/v3/common";_d "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_ag "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_ef "io";);func (_ebc *Encoder )code0 (_edd *codingContext ,_cfa uint32 ,_dfg uint16 ,_ecc byte ){if _edd .mps (_cfa )==0{_ebc .codeMPS (_edd ,_cfa ,_dfg ,_ecc );
}else {_ebc .codeLPS (_edd ,_cfa ,_dfg ,_ecc );};};const _dea =0x9b25;var _efad =[]state {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
var _f =[]intEncRangeS {{0,3,0,2,0,2},{-1,-1,9,4,0,0},{-3,-2,5,3,2,1},{4,19,2,3,4,4},{-19,-4,3,3,4,4},{20,83,6,4,20,6},{-83,-20,7,4,20,6},{84,339,14,5,84,8},{-339,-84,15,5,84,8},{340,4435,30,6,340,12},{-4435,-340,31,6,340,12},{4436,2000000000,62,6,4436,32},{-2000000000,-4436,63,6,4436,32}};
func _acg (_fg int )*codingContext {return &codingContext {_eb :make ([]byte ,_fg ),_dac :make ([]byte ,_fg )};};type codingContext struct{_eb []byte ;_dac []byte ;};func (_fb *codingContext )flipMps (_gd uint32 ){_fb ._dac [_gd ]=1-_fb ._dac [_gd ]};func (_afad *Encoder )code1 (_bdg *codingContext ,_bdd uint32 ,_gec uint16 ,_cggf byte ){if _bdg .mps (_bdd )==1{_afad .codeMPS (_bdg ,_bdd ,_gec ,_cggf );
}else {_afad .codeLPS (_bdg ,_bdd ,_gec ,_cggf );};};type intEncRangeS struct{_da ,_aef int ;_de ,_ac uint8 ;_ad uint16 ;_be uint8 ;};const (IAAI Class =iota ;IADH ;IADS ;IADT ;IADW ;IAEX ;IAFS ;IAIT ;IARDH ;IARDW ;IARDX ;IARDY ;IARI ;);func (_ee *Encoder )codeLPS (_fbga *codingContext ,_ffd uint32 ,_bge uint16 ,_aec byte ){_ee ._ebe -=_bge ;
if _ee ._ebe < _bge {_ee ._ab +=uint32 (_bge );}else {_ee ._ebe =_bge ;};if _efad [_aec ]._bdge ==1{_fbga .flipMps (_ffd );};_fbga ._eb [_ffd ]=_efad [_aec ]._aab ;_ee .renormalize ();};func (_gbbf *Encoder )encodeInteger (_gffa Class ,_gdbb int )error {const _fdg ="E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072";
if _gdbb > 2000000000||_gdbb < -2000000000{return _ag .Errorf (_fdg ,"\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",_gdbb );
};_ebcb :=_gbbf ._cb [_gffa ];_fab :=uint32 (1);var _agga int ;for ;;_agga ++{if _f [_agga ]._da <=_gdbb &&_f [_agga ]._aef >=_gdbb {break ;};};if _gdbb < 0{_gdbb =-_gdbb ;};_gdbb -=int (_f [_agga ]._ad );_eaa :=_f [_agga ]._de ;for _decg :=uint8 (0);_decg < _f [_agga ]._ac ;
_decg ++{_gea :=_eaa &1;if _eba :=_gbbf .encodeBit (_ebcb ,_fab ,_gea );_eba !=nil {return _ag .Wrap (_eba ,_fdg ,"");};_eaa >>=1;if _fab &0x100> 0{_fab =(((_fab <<1)|uint32 (_gea ))&0x1ff)|0x100;}else {_fab =(_fab <<1)|uint32 (_gea );};};_gdbb <<=32-_f [_agga ]._be ;
for _abg :=uint8 (0);_abg < _f [_agga ]._be ;_abg ++{_cgae :=uint8 ((uint32 (_gdbb )&0x80000000)>>31);if _gfgc :=_gbbf .encodeBit (_ebcb ,_fab ,_cgae );_gfgc !=nil {return _ag .Wrap (_gfgc ,_fdg ,"\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064");
};_gdbb <<=1;if _fab &0x100!=0{_fab =(((_fab <<1)|uint32 (_cgae ))&0x1ff)|0x100;}else {_fab =(_fab <<1)|uint32 (_cgae );};};return nil ;};func (_aae *Encoder )byteOut (){if _aae ._bc ==0xff{_aae .rBlock ();return ;};if _aae ._ab < 0x8000000{_aae .lBlock ();
return ;};_aae ._bc ++;if _aae ._bc !=0xff{_aae .lBlock ();return ;};_aae ._ab &=0x7ffffff;_aae .rBlock ();};func (_dec *Encoder )EncodeOOB (proc Class )(_adg error ){_a .Log .Trace ("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027",proc );
if _adg =_dec .encodeOOB (proc );_adg !=nil {return _ag .Wrap (_adg ,"\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB","");};return nil ;};func (_dda *Encoder )flush (){_dda .setBits ();_dda ._ab <<=_dda ._efc ;_dda .byteOut ();_dda ._ab <<=_dda ._efc ;_dda .byteOut ();
_dda .emit ();if _dda ._bc !=0xff{_dda ._df ++;_dda ._bc =0xff;_dda .emit ();};_dda ._df ++;_dda ._bc =0xac;_dda ._df ++;_dda .emit ();};func (_aefd *Encoder )dataSize ()int {return _gfd *len (_aefd ._ff )+_aefd ._c };type Class int ;func (_edbb *Encoder )renormalize (){for {_edbb ._ebe <<=1;
_edbb ._ab <<=1;_edbb ._efc --;if _edbb ._efc ==0{_edbb .byteOut ();};if (_edbb ._ebe &0x8000)!=0{break ;};};};func (_gbb *Encoder )Reset (){_gbb ._ebe =0x8000;_gbb ._ab =0;_gbb ._efc =12;_gbb ._df =-1;_gbb ._bc =0;_gbb ._aa =nil ;_gbb ._ga =_acg (_ggg );
};func (_cd *Encoder )EncodeInteger (proc Class ,value int )(_ca error ){_a .Log .Trace ("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027",value ,proc );
if _ca =_cd .encodeInteger (proc ,value );_ca !=nil {return _ag .Wrap (_ca ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072","");};return nil ;};var _ _ef .WriterTo =&Encoder {};type Encoder struct{_ab uint32 ;_ebe uint16 ;
_efc ,_bc uint8 ;_df int ;_ea int ;_ff [][]byte ;_fc []byte ;_c int ;_ga *codingContext ;_cb [13]*codingContext ;_aa *codingContext ;};func (_cga *Encoder )WriteTo (w _ef .Writer )(int64 ,error ){const _ebg ="\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f";
var _ec int64 ;for _aad ,_eag :=range _cga ._ff {_gc ,_agg :=w .Write (_eag );if _agg !=nil {return 0,_ag .Wrapf (_agg ,_ebg ,"\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b",_aad );
};_ec +=int64 (_gc );};_cga ._fc =_cga ._fc [:_cga ._c ];_ffb ,_gfa :=w .Write (_cga ._fc );if _gfa !=nil {return 0,_ag .Wrap (_gfa ,_ebg ,"\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073");};_ec +=int64 (_ffb );return _ec ,nil ;
};const (_ggg =65536;_gfd =20*1024;);func (_fbg *Encoder )Flush (){_fbg ._c =0;_fbg ._ff =nil ;_fbg ._df =-1};func (_ecd *Encoder )rBlock (){if _ecd ._df >=0{_ecd .emit ();};_ecd ._df ++;_ecd ._bc =uint8 (_ecd ._ab >>20);_ecd ._ab &=0xfffff;_ecd ._efc =7;
};func (_ede *Encoder )setBits (){_adef :=_ede ._ab +uint32 (_ede ._ebe );_ede ._ab |=0xffff;if _ede ._ab >=_adef {_ede ._ab -=0x8000;};};func (_g *codingContext )mps (_efe uint32 )int {return int (_g ._dac [_efe ])};func (_cbe *Encoder )codeMPS (_cfg *codingContext ,_ebf uint32 ,_ce uint16 ,_cdg byte ){_cbe ._ebe -=_ce ;
if _cbe ._ebe &0x8000!=0{_cbe ._ab +=uint32 (_ce );return ;};if _cbe ._ebe < _ce {_cbe ._ebe =_ce ;}else {_cbe ._ab +=uint32 (_ce );};_cfg ._eb [_ebf ]=_efad [_cdg ]._ddf ;_cbe .renormalize ();};func (_ae Class )String ()string {switch _ae {case IAAI :return "\u0049\u0041\u0041\u0049";
case IADH :return "\u0049\u0041\u0044\u0048";case IADS :return "\u0049\u0041\u0044\u0053";case IADT :return "\u0049\u0041\u0044\u0054";case IADW :return "\u0049\u0041\u0044\u0057";case IAEX :return "\u0049\u0041\u0045\u0058";case IAFS :return "\u0049\u0041\u0046\u0053";
case IAIT :return "\u0049\u0041\u0049\u0054";case IARDH :return "\u0049\u0041\u0052D\u0048";case IARDW :return "\u0049\u0041\u0052D\u0057";case IARDX :return "\u0049\u0041\u0052D\u0058";case IARDY :return "\u0049\u0041\u0052D\u0059";case IARI :return "\u0049\u0041\u0052\u0049";
default:return "\u0055N\u004b\u004e\u004f\u0057\u004e";};};func (_ecf *Encoder )encodeIAID (_edb ,_afg int )error {if _ecf ._aa ==nil {_ecf ._aa =_acg (1<<uint (_edb ));};_bga :=uint32 (1<<uint32 (_edb +1))-1;_afg <<=uint (32-_edb );_fbdc :=uint32 (1);
for _aaf :=0;_aaf < _edb ;_aaf ++{_agd :=_fbdc &_bga ;_gaf :=uint8 ((uint32 (_afg )&0x80000000)>>31);if _cfaf :=_ecf .encodeBit (_ecf ._aa ,_agd ,_gaf );_cfaf !=nil {return _cfaf ;};_fbdc =(_fbdc <<1)|uint32 (_gaf );_afg <<=1;};return nil ;};func (_ggd *Encoder )encodeBit (_eeg *codingContext ,_dge uint32 ,_fca uint8 )error {const _gab ="\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074";
_ggd ._ea ++;if _dge >=uint32 (len (_eeg ._eb )){return _ag .Errorf (_gab ,"\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027",_dge );
};_bfa :=_eeg ._eb [_dge ];_ffdc :=_eeg .mps (_dge );_ffa :=_efad [_bfa ]._aefg ;_a .Log .Trace ("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064",_ggd ._ea ,_fca ,_bfa ,_ffdc ,_ffa ,_ggd ._ebe ,_ggd ._ab ,_ggd ._efc ,_ggd ._bc ,_ggd ._df );
if _fca ==0{_ggd .code0 (_eeg ,_dge ,_ffa ,_bfa );}else {_ggd .code1 (_eeg ,_dge ,_ffa ,_bfa );};return nil ;};func New ()*Encoder {_db :=&Encoder {};_db .Init ();return _db };func (_gb *Encoder )EncodeBitmap (bm *_d .Bitmap ,duplicateLineRemoval bool )error {_a .Log .Trace ("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s",bm .Width ,bm .Height ,bm );
var (_gdf ,_dfe uint8 ;_ddb ,_dg ,_bf uint16 ;_bce ,_fbd ,_ba byte ;_fe ,_gdb ,_cg int ;_afa ,_baa []byte ;);for _bd :=0;_bd < bm .Height ;_bd ++{_bce ,_fbd =0,0;if _bd >=2{_bce =bm .Data [(_bd -2)*bm .RowStride ];};if _bd >=1{_fbd =bm .Data [(_bd -1)*bm .RowStride ];
if duplicateLineRemoval {_gdb =_bd *bm .RowStride ;_afa =bm .Data [_gdb :_gdb +bm .RowStride ];_cg =(_bd -1)*bm .RowStride ;_baa =bm .Data [_cg :_cg +bm .RowStride ];if _b .Equal (_afa ,_baa ){_dfe =_gdf ^1;_gdf =1;}else {_dfe =_gdf ;_gdf =0;};};};if duplicateLineRemoval {if _acgb :=_gb .encodeBit (_gb ._ga ,_dea ,_dfe );
_acgb !=nil {return _acgb ;};if _gdf !=0{continue ;};};_ba =bm .Data [_bd *bm .RowStride ];_ddb =uint16 (_bce >>5);_dg =uint16 (_fbd >>4);_bce <<=3;_fbd <<=4;_bf =0;for _fe =0;_fe < bm .Width ;_fe ++{_gbf :=uint32 (_ddb <<11|_dg <<4|_bf );_gff :=(_ba &0x80)>>7;
_bag :=_gb .encodeBit (_gb ._ga ,_gbf ,_gff );if _bag !=nil {return _bag ;};_ddb <<=1;_dg <<=1;_bf <<=1;_ddb |=uint16 ((_bce &0x80)>>7);_dg |=uint16 ((_fbd &0x80)>>7);_bf |=uint16 (_gff );_cgg :=_fe %8;_gg :=_fe /8+1;if _cgg ==4&&_bd >=2{_bce =0;if _gg < bm .RowStride {_bce =bm .Data [(_bd -2)*bm .RowStride +_gg ];
};}else {_bce <<=1;};if _cgg ==3&&_bd >=1{_fbd =0;if _gg < bm .RowStride {_fbd =bm .Data [(_bd -1)*bm .RowStride +_gg ];};}else {_fbd <<=1;};if _cgg ==7{_ba =0;if _gg < bm .RowStride {_ba =bm .Data [_bd *bm .RowStride +_gg ];};}else {_ba <<=1;};_ddb &=31;
_dg &=127;_bf &=15;};};return nil ;};func (_efd *Encoder )Final (){_efd .flush ()};func (_fcag *Encoder )lBlock (){if _fcag ._df >=0{_fcag .emit ();};_fcag ._df ++;_fcag ._bc =uint8 (_fcag ._ab >>19);_fcag ._ab &=0x7ffff;_fcag ._efc =8;};func (_aga *Encoder )EncodeIAID (symbolCodeLength ,value int )(_bff error ){_a .Log .Trace ("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",symbolCodeLength ,value );
if _bff =_aga .encodeIAID (symbolCodeLength ,value );_bff !=nil {return _ag .Wrap (_bff ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044","");};return nil ;};func (_cgc *Encoder )encodeOOB (_ade Class )error {_ecg :=_cgc ._cb [_ade ];_bee :=_cgc .encodeBit (_ecg ,1,1);
if _bee !=nil {return _bee ;};_bee =_cgc .encodeBit (_ecg ,3,0);if _bee !=nil {return _bee ;};_bee =_cgc .encodeBit (_ecg ,6,0);if _bee !=nil {return _bee ;};_bee =_cgc .encodeBit (_ecg ,12,0);if _bee !=nil {return _bee ;};return nil ;};func (_agae *Encoder )emit (){if _agae ._c ==_gfd {_agae ._ff =append (_agae ._ff ,_agae ._fc );
_agae ._fc =make ([]byte ,_gfd );_agae ._c =0;};_agae ._fc [_agae ._c ]=_agae ._bc ;_agae ._c ++;};func (_dd *Encoder )Init (){_dd ._ga =_acg (_ggg );_dd ._ebe =0x8000;_dd ._ab =0;_dd ._efc =12;_dd ._df =-1;_dd ._bc =0;_dd ._c =0;_dd ._fc =make ([]byte ,_gfd );
for _gf :=0;_gf < len (_dd ._cb );_gf ++{_dd ._cb [_gf ]=_acg (512);};_dd ._aa =nil ;};type state struct{_aefg uint16 ;_ddf ,_aab uint8 ;_bdge uint8 ;};func (_af *Encoder )DataSize ()int {return _af .dataSize ()};func (_gda *Encoder )Refine (iTemp ,iTarget *_d .Bitmap ,ox ,oy int )error {for _gag :=0;
_gag < iTarget .Height ;_gag ++{var _abe int ;_efa :=_gag +oy ;var (_acb ,_ed ,_cf ,_bg ,_afae uint16 ;_gfg ,_gge ,_agag ,_fd ,_bgb byte ;);if _efa >=1&&(_efa -1)< iTemp .Height {_gfg =iTemp .Data [(_efa -1)*iTemp .RowStride ];};if _efa >=0&&_efa < iTemp .Height {_gge =iTemp .Data [_efa *iTemp .RowStride ];
};if _efa >=-1&&_efa +1< iTemp .Height {_agag =iTemp .Data [(_efa +1)*iTemp .RowStride ];};if _gag >=1{_fd =iTarget .Data [(_gag -1)*iTarget .RowStride ];};_bgb =iTarget .Data [_gag *iTarget .RowStride ];_cge :=uint (6+ox );_acb =uint16 (_gfg >>_cge );
_ed =uint16 (_gge >>_cge );_cf =uint16 (_agag >>_cge );_bg =uint16 (_fd >>6);_ge :=uint (2-ox );_gfg <<=_ge ;_gge <<=_ge ;_agag <<=_ge ;_fd <<=2;for _abe =0;_abe < iTarget .Width ;_abe ++{_fa :=(_acb <<10)|(_ed <<7)|(_cf <<4)|(_bg <<1)|_afae ;_dga :=_bgb >>7;
_bfd :=_gda .encodeBit (_gda ._ga ,uint32 (_fa ),_dga );if _bfd !=nil {return _bfd ;};_acb <<=1;_ed <<=1;_cf <<=1;_bg <<=1;_acb |=uint16 (_gfg >>7);_ed |=uint16 (_gge >>7);_cf |=uint16 (_agag >>7);_bg |=uint16 (_fd >>7);_afae =uint16 (_dga );_efb :=_abe %8;
_fac :=_abe /8+1;if _efb ==5+ox {_gfg ,_gge ,_agag =0,0,0;if _fac < iTemp .RowStride &&_efa >=1&&(_efa -1)< iTemp .Height {_gfg =iTemp .Data [(_efa -1)*iTemp .RowStride +_fac ];};if _fac < iTemp .RowStride &&_efa >=0&&_efa < iTemp .Height {_gge =iTemp .Data [_efa *iTemp .RowStride +_fac ];
};if _fac < iTemp .RowStride &&_efa >=-1&&(_efa +1)< iTemp .Height {_agag =iTemp .Data [(_efa +1)*iTemp .RowStride +_fac ];};}else {_gfg <<=1;_gge <<=1;_agag <<=1;};if _efb ==5&&_gag >=1{_fd =0;if _fac < iTarget .RowStride {_fd =iTarget .Data [(_gag -1)*iTarget .RowStride +_fac ];
};}else {_fd <<=1;};if _efb ==7{_bgb =0;if _fac < iTarget .RowStride {_bgb =iTarget .Data [_gag *iTarget .RowStride +_fac ];};}else {_bgb <<=1;};_acb &=7;_ed &=7;_cf &=7;_bg &=7;};};return nil ;};