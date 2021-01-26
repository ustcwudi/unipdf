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

package arithmetic ;import (_f "fmt";_fb "github.com/unidoc/unipdf/v3/common";_a "github.com/unidoc/unipdf/v3/internal/bitwise";_fa "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_g "io";_d "strings";);func NewStats (contextSize int32 ,index int32 )*DecoderStats {return &DecoderStats {_ade :index ,_baa :contextSize ,_aa :make ([]byte ,contextSize ),_cbc :make ([]byte ,contextSize )};};func (_dce *Decoder )readByte ()error {if _dce ._gc .StreamPosition ()> _dce ._ba {if _ ,_fab :=_dce ._gc .Seek (-1,_g .SeekCurrent );_fab !=nil {return _fab ;};};_ded ,_fdf :=_dce ._gc .ReadByte ();if _fdf !=nil {return _fdf ;};_dce ._b =_ded ;if _dce ._b ==0xFF{_aef ,_cgg :=_dce ._gc .ReadByte ();if _cgg !=nil {return _cgg ;};if _aef > 0x8F{_dce ._ed +=0xFF00;_dce ._fg =8;if _ ,_gd :=_dce ._gc .Seek (-2,_g .SeekCurrent );_gd !=nil {return _gd ;};}else {_dce ._ed +=uint64 (_aef )<<9;_dce ._fg =7;};}else {_ded ,_fdf =_dce ._gc .ReadByte ();if _fdf !=nil {return _fdf ;};_dce ._b =_ded ;_dce ._ed +=uint64 (_dce ._b )<<8;_dce ._fg =8;};_dce ._ed &=0xFFFFFFFFFF;return nil ;};func (_fbf *Decoder )mpsExchange (_add *DecoderStats ,_ffb int32 )int {_gcg :=_add ._cbc [_add ._ade ];if _fbf ._eg < _eb [_ffb ][0]{if _eb [_ffb ][3]==1{_add .toggleMps ();};_add .setEntry (int (_eb [_ffb ][2]));return int (1-_gcg );};_add .setEntry (int (_eb [_ffb ][1]));return int (_gcg );};func (_de *Decoder )DecodeInt (stats *DecoderStats )(int32 ,error ){var (_dc ,_ag int32 ;_fdd ,_df ,_ae int ;_cf error ;);if stats ==nil {stats =NewStats (512,1);};_de ._fd =1;_df ,_cf =_de .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};_fdd ,_cf =_de .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};if _fdd ==1{_fdd ,_cf =_de .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};if _fdd ==1{_fdd ,_cf =_de .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};if _fdd ==1{_fdd ,_cf =_de .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};if _fdd ==1{_fdd ,_cf =_de .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};if _fdd ==1{_ae =32;_ag =4436;}else {_ae =12;_ag =340;};}else {_ae =8;_ag =84;};}else {_ae =6;_ag =20;};}else {_ae =4;_ag =4;};}else {_ae =2;_ag =0;};for _dg :=0;_dg < _ae ;_dg ++{_fdd ,_cf =_de .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};_dc =(_dc <<1)|int32 (_fdd );};_dc +=_ag ;if _df ==0{return _dc ,nil ;}else if _df ==1&&_dc > 0{return -_dc ,nil ;};return 0,_fa .ErrOOB ;};func (_gdf *DecoderStats )Overwrite (dNew *DecoderStats ){for _agd :=0;_agd < len (_gdf ._aa );_agd ++{_gdf ._aa [_agd ]=dNew ._aa [_agd ];_gdf ._cbc [_agd ]=dNew ._cbc [_agd ];};};var (_eb =[][4]uint32 {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};);func (_eda *Decoder )init ()error {_eda ._ba =_eda ._gc .StreamPosition ();_gcc ,_cg :=_eda ._gc .ReadByte ();if _cg !=nil {_fb .Log .Debug ("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076",_cg );return _cg ;};_eda ._b =_gcc ;_eda ._ed =uint64 (_gcc )<<16;if _cg =_eda .readByte ();_cg !=nil {return _cg ;};_eda ._ed <<=7;_eda ._fg -=7;_eda ._eg =0x8000;_eda ._gb ++;return nil ;};func (_gcb *DecoderStats )toggleMps (){_gcb ._cbc [_gcb ._ade ]^=1};func (_adab *DecoderStats )setEntry (_acf int ){_gdd :=byte (_acf &0x7f);_adab ._aa [_adab ._ade ]=_gdd ;};func (_gf *DecoderStats )Reset (){for _bc :=0;_bc < len (_gf ._aa );_bc ++{_gf ._aa [_bc ]=0;_gf ._cbc [_bc ]=0;};};func (_ada *DecoderStats )String ()string {_bfd :=&_d .Builder {};_bfd .WriteString (_f .Sprintf ("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a",len (_ada ._aa )));for _bcf ,_dgb :=range _ada ._aa {if _dgb !=0{_bfd .WriteString (_f .Sprintf ("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a",_bcf ,_dgb ));};};return _bfd .String ();};func (_fc *Decoder )renormalize ()error {for {if _fc ._fg ==0{if _fae :=_fc .readByte ();_fae !=nil {return _fae ;};};_fc ._eg <<=1;_fc ._ed <<=1;_fc ._fg --;if (_fc ._eg &0x8000)!=0{break ;};};_fc ._ed &=0xffffffff;return nil ;};func New (r _a .StreamReader )(*Decoder ,error ){_db :=&Decoder {_gc :r ,ContextSize :[]uint32 {16,13,10,10},ReferedToContextSize :[]uint32 {13,10}};if _c :=_db .init ();_c !=nil {return nil ,_c ;};return _db ,nil ;};func (_gccc *DecoderStats )cx ()byte {return _gccc ._aa [_gccc ._ade ]};func (_ca *DecoderStats )getMps ()byte {return _ca ._cbc [_ca ._ade ]};func (_cea *Decoder )decodeIntBit (_cd *DecoderStats )(int ,error ){_cd .SetIndex (int32 (_cea ._fd ));_edge ,_gcf :=_cea .DecodeBit (_cd );if _gcf !=nil {_fb .Log .Debug ("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076",_gcf );return _edge ,_gcf ;};if _cea ._fd < 256{_cea ._fd =((_cea ._fd <<uint64 (1))|int64 (_edge ))&0x1ff;}else {_cea ._fd =(((_cea ._fd <<uint64 (1)|int64 (_edge ))&511)|256)&0x1ff;};return _edge ,nil ;};type DecoderStats struct{_ade int32 ;_baa int32 ;_aa []byte ;_cbc []byte ;};type Decoder struct{ContextSize []uint32 ;ReferedToContextSize []uint32 ;_gc _a .StreamReader ;_b uint8 ;_ed uint64 ;_eg uint32 ;_fd int64 ;_fg int32 ;_gb int32 ;_ba int64 ;};func (_ce *Decoder )DecodeIAID (codeLen uint64 ,stats *DecoderStats )(int64 ,error ){_ce ._fd =1;var _afd uint64 ;for _afd =0;_afd < codeLen ;_afd ++{stats .SetIndex (int32 (_ce ._fd ));_ebc ,_ab :=_ce .DecodeBit (stats );if _ab !=nil {return 0,_ab ;};_ce ._fd =(_ce ._fd <<1)|int64 (_ebc );};_gbe :=_ce ._fd -(1<<codeLen );return _gbe ,nil ;};func (_af *Decoder )DecodeBit (stats *DecoderStats )(int ,error ){var (_bf int ;_ad =_eb [stats .cx ()][0];_ff =int32 (stats .cx ()););defer func (){_af ._gb ++}();_af ._eg -=_ad ;if (_af ._ed >>16)< uint64 (_ad ){_bf =_af .lpsExchange (stats ,_ff ,_ad );if _edg :=_af .renormalize ();_edg !=nil {return 0,_edg ;};}else {_af ._ed -=uint64 (_ad )<<16;if (_af ._eg &0x8000)==0{_bf =_af .mpsExchange (stats ,_ff );if _dd :=_af .renormalize ();_dd !=nil {return 0,_dd ;};}else {_bf =int (stats .getMps ());};};return _bf ,nil ;};func (_cb *Decoder )lpsExchange (_cga *DecoderStats ,_gde int32 ,_gbf uint32 )int {_ee :=_cga .getMps ();if _cb ._eg < _gbf {_cga .setEntry (int (_eb [_gde ][1]));_cb ._eg =_gbf ;return int (_ee );};if _eb [_gde ][3]==1{_cga .toggleMps ();};_cga .setEntry (int (_eb [_gde ][2]));_cb ._eg =_gbf ;return int (1-_ee );};func (_afg *DecoderStats )SetIndex (index int32 ){_afg ._ade =index };func (_ac *DecoderStats )Copy ()*DecoderStats {_fff :=&DecoderStats {_baa :_ac ._baa ,_aa :make ([]byte ,_ac ._baa )};for _ea :=0;_ea < len (_ac ._aa );_ea ++{_fff ._aa [_ea ]=_ac ._aa [_ea ];};return _fff ;};