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

package errors ;import (_a "fmt";_fg "golang.org/x/xerrors";);func Wrap (err error ,processName ,message string )error {if _bgf ,_cb :=err .(*processError );_cb {_bgf ._ae ="";};_d :=_gf (message ,processName );_d ._c =err ;return _d ;};func Errorf (processName ,message string ,arguments ...interface{})error {return _gf (_a .Sprintf (message ,arguments ...),processName );
};func (_fgg *processError )Unwrap ()error {return _fgg ._c };func (_g *processError )Error ()string {var _ac string ;if _g ._ae !=""{_ac =_g ._ae ;};_ac +="\u0050r\u006f\u0063\u0065\u0073\u0073\u003a "+_g ._b ;if _g ._e !=""{_ac +="\u0020\u004d\u0065\u0073\u0073\u0061\u0067\u0065\u003a\u0020"+_g ._e ;
};if _g ._c !=nil {_ac +="\u002e\u0020"+_g ._c .Error ();};return _ac ;};func Wrapf (err error ,processName ,message string ,arguments ...interface{})error {if _df ,_db :=err .(*processError );_db {_df ._ae ="";};_dc :=_gf (_a .Sprintf (message ,arguments ...),processName );
_dc ._c =err ;return _dc ;};type processError struct{_ae string ;_b string ;_e string ;_c error ;};func Error (processName ,message string )error {return _gf (message ,processName )};func _gf (_bg ,_gg string )*processError {return &processError {_ae :"\u005b\u0055\u006e\u0069\u0050\u0044\u0046\u005d",_e :_bg ,_b :_gg };
};var _ _fg .Wrapper =(*processError )(nil );