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

package decoder ;import (_e "github.com/unidoc/unipdf/v3/internal/bitwise";_a "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_d "github.com/unidoc/unipdf/v3/internal/jbig2/document";_gb "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_b "image";
);func (_db *Decoder )decodePage (_eg int )([]byte ,error ){const _fd ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";if _eg < 0{return nil ,_gb .Errorf (_fd ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_eg );
};if _eg > int (_db ._ag .NumberOfPages ){return nil ,_gb .Errorf (_fd ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_eg );
};_ce ,_ab :=_db ._ag .GetPage (_eg );if _ab !=nil {return nil ,_gb .Wrap (_ab ,_fd ,"");};_ee ,_ab :=_ce .GetBitmap ();if _ab !=nil {return nil ,_gb .Wrap (_ab ,_fd ,"");};_ee .InverseData ();if !_db ._gbf .UnpaddedData {return _ee .Data ,nil ;};return _ee .GetUnpaddedData ();
};func (_bg *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _bg .decodePage (pageNumber )};func (_f *Decoder )DecodePageImage (pageNumber int )(_b .Image ,error ){const _ga ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";
_c ,_gg :=_f .decodePageImage (pageNumber );if _gg !=nil {return nil ,_gb .Wrap (_gg ,_ga ,"");};return _c ,nil ;};func Decode (input []byte ,parameters Parameters ,globals *_d .Globals )(*Decoder ,error ){_acg :=_e .NewReader (input );_abe ,_bc :=_d .DecodeDocument (_acg ,globals );
if _bc !=nil {return nil ,_bc ;};return &Decoder {_ba :_acg ,_ag :_abe ,_gbf :parameters },nil ;};func (_cb *Decoder )DecodeNextPage ()([]byte ,error ){_cb ._ac ++;_acd :=_cb ._ac ;return _cb .decodePage (_acd );};func (_bb *Decoder )decodePageImage (_bgb int )(_b .Image ,error ){const _cea ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";
if _bgb < 0{return nil ,_gb .Errorf (_cea ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_bgb );};if _bgb > int (_bb ._ag .NumberOfPages ){return nil ,_gb .Errorf (_cea ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_bgb );
};_da ,_ggd :=_bb ._ag .GetPage (_bgb );if _ggd !=nil {return nil ,_gb .Wrap (_ggd ,_cea ,"");};_ea ,_ggd :=_da .GetBitmap ();if _ggd !=nil {return nil ,_gb .Wrap (_ggd ,_cea ,"");};_ea .InverseData ();return _ea .ToImage (),nil ;};type Parameters struct{UnpaddedData bool ;
Color _a .Color ;};type Decoder struct{_ba _e .StreamReader ;_ag *_d .Document ;_ac int ;_gbf Parameters ;};func (_de *Decoder )PageNumber ()(int ,error ){const _cg ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";
if _de ._ag ==nil {return 0,_gb .Error (_cg ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");};return int (_de ._ag .NumberOfPages ),nil ;};