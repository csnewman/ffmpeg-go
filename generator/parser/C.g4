grammar C;

unit
    : topLevel* EOF
    ;

topLevel
    : Ifndef topLevel* Endif #topLevelIfndef
    | Ifdir topLevel* Endif #topLevelIfdir
    | Define #topLevelDefine
    | Undef #topLevelUndef
    | Include #topLevelInclude
    | Pragma #topLevelPragma
    | Enum (itemname=Identifier)? LBrace enumBody Comma? RBrace Semi #topLevelEnum
    | AttributeDeprecated? Typedef Struct structname=Identifier LBrace structItem* RBrace itemname=Identifier Semi #topLevelStructTypeDef
//    | AttributeDeprecated? Typedef Struct Identifier Identifier Semi #topLevelStructTypeDefOpaque
    | AttributeDeprecated? Struct AttributeDeprecated? itemname=Identifier LBrace structItem* RBrace Semi #topLevelStruct
//    | AttributeDeprecated? Struct AttributeDeprecated? Identifier Semi #topLevelStructOpaque
    | AttributeDeprecated? Typedef type itemname=Identifier Semi #topLevelTypeDefSimple
    | AttributeDeprecated? Typedef type LParen Star itemname=Identifier RParen LParen params RParen Semi #topLevelTypeDefFunc
    | AttributeDeprecated? type itemname=Identifier LParen params RParen Semi #topLevelFunction
    ;

enumBody
    : Define enumBody #enumBodyDefine
    | Ifdir (enumItem Comma)+ Endif enumBody #enumBodyIfdir
    | enumItem Comma enumBody #enumBodyItem
    | enumItem #enumBodyEnd
    ;

enumItem
    : Identifier #enumItemNoValue
    | Identifier Assign constant #enumItemConstant
    ;

structItem
    : AttributeDeprecated? type (Identifier Comma)* Identifier Semi #structItemMember
    | AttributeDeprecated? type Identifier arraySize Semi #structItemArray
    | AttributeDeprecated? type LParen Star Identifier RParen arraySize Semi #structItemArrayPtr
    | Ifdir structItem* Endif #structItemIfdir
    | Define #structItemDefine
    | type LParen Star Identifier RParen LParen params RParen Semi #structItemFunction
    ;

constant
    : (min=Minus)? Constant
    ;

params
    : Void #paramsVoid
    | (param Comma)* param #paramsValues
    ;

param
    : type Identifier #paramSimple
    | type Identifier arraySize #paramArray
    | type #paramAnon
    | type LParen Star Identifier RParen LParen params RParen #paramFunction
    ;

arraySize
    : arraySizeDim+
    ;

arraySizeDim
    : LBracket Identifier RBracket #arraySizeDimIden
    | LBracket Constant RBracket #arraySizeDimConst
    ;

type
    : typeInner #typePass
    | type Star #typePtr
    | type Star Const #typePtrConst
    ;

typeInner
    : Identifier #typeInnerIdent
    | Enum Identifier #typeInnerEnum
    | Struct Identifier #typeInnerStruct
    | Unsigned Identifier? #typeInnerUnsigned
    | Void #typeInnerVoid
    | Const typeInner #typeInnerConst
    ;

Assign: '=';
Star: '*';
Minus: '-';

Comma: ',';
Semi: ';';

LParen: '(';
RParen: ')';
LBracket: '[';
RBracket: ']';
LBrace: '{';
RBrace: '}';

Const: 'const';
Unsigned: 'unsigned';

Void: 'void';
Enum: 'enum';
Struct: 'struct';

Typedef: 'typedef';

AttributeDeprecated: 'attribute_deprecated';


Ifndef
//    : '#ifndef' WhitespaceFrag IdentifierFrag WhitespaceFrag? NewlineFrag
    : '#ifndef' ~[\r\n]* NewlineFrag
    ;


Ifdir
//    : '#ifndef' WhitespaceFrag IdentifierFrag WhitespaceFrag? NewlineFrag
    : '#if' ~[\r\n]* NewlineFrag
    ;

Endif
    : '#endif' ~[\r\n]* NewlineFrag
    ;

Define
//    : '#define' WhitespaceFrag IdentifierFrag WhitespaceFrag? NewlineFrag
//    : '#define' ~[\r\n]* NewlineFrag
    : '#define' (~[\r\n]* (('\\' NewlineFrag) | ( '/*' .*? '*/')))* ~[\r\n]* NewlineFrag
    ;

Undef
    : '#undef' ~[\r\n]* NewlineFrag
    ;

Pragma
    : '#pragma' ~[\r\n]* NewlineFrag
    ;

Include
    : '#include' WhitespaceFrag '<' (IdentifierFrag '/')* IdentifierFrag '.h>' WhitespaceFrag? NewlineFrag
    | '#include' WhitespaceFrag '"' (IdentifierFrag '/')* IdentifierFrag '.h"' WhitespaceFrag? NewlineFrag
    ;

Identifier
    : IdentifierFrag
    ;

fragment
IdentifierFrag
    :   IdentifierNondigit
        (   IdentifierNondigit
        |   Digit
        )*
    ;

fragment
IdentifierNondigit
    : Nondigit
    | UniversalCharacterName
    ;

fragment
Nondigit
    : [a-zA-Z_]
    ;

fragment
Digit
    : [0-9]
    ;

fragment
UniversalCharacterName
    :   '\\u' HexQuad
    |   '\\U' HexQuad HexQuad
    ;

fragment
HexQuad
    :   HexadecimalDigit HexadecimalDigit HexadecimalDigit HexadecimalDigit
    ;

Whitespace
    : WhitespaceFrag -> channel(HIDDEN)
    ;

fragment
WhitespaceFrag
    : [ \t]+
    ;

Newline
    : NewlineFrag -> channel(HIDDEN)
    ;

Constant
    :   IntegerConstant
    |   FloatingConstant
    |   CharacterConstant
    ;

fragment
IntegerConstant
    :   DecimalConstant IntegerSuffix?
    |   OctalConstant IntegerSuffix?
    |   HexadecimalConstant IntegerSuffix?
    |	BinaryConstant
    ;

fragment
BinaryConstant
	:	'0' [bB] [0-1]+
	;

fragment
DecimalConstant
    :   NonzeroDigit Digit*
    ;

fragment
OctalConstant
    :   '0' OctalDigit*
    ;

fragment
HexadecimalConstant
    :   HexadecimalPrefix HexadecimalDigit+
    ;

fragment
HexadecimalPrefix
    :   '0' [xX]
    ;

fragment
NonzeroDigit
    :   [1-9]
    ;

fragment
OctalDigit
    :   [0-7]
    ;

fragment
HexadecimalDigit
    :   [0-9a-fA-F]
    ;

fragment
IntegerSuffix
    :   UnsignedSuffix LongSuffix?
    |   UnsignedSuffix LongLongSuffix
    |   LongSuffix UnsignedSuffix?
    |   LongLongSuffix UnsignedSuffix?
    ;

fragment
UnsignedSuffix
    :   [uU]
    ;

fragment
LongSuffix
    :   [lL]
    ;

fragment
LongLongSuffix
    :   'll' | 'LL'
    ;

fragment
FloatingConstant
    :   DecimalFloatingConstant
    |   HexadecimalFloatingConstant
    ;

fragment
DecimalFloatingConstant
    :   FractionalConstant ExponentPart? FloatingSuffix?
    |   DigitSequence ExponentPart FloatingSuffix?
    ;

fragment
HexadecimalFloatingConstant
    :   HexadecimalPrefix (HexadecimalFractionalConstant | HexadecimalDigitSequence) BinaryExponentPart FloatingSuffix?
    ;

fragment
FractionalConstant
    :   DigitSequence? '.' DigitSequence
    |   DigitSequence '.'
    ;

fragment
ExponentPart
    :   [eE] Sign? DigitSequence
    ;

fragment
Sign
    :   [+-]
    ;

DigitSequence
    :   Digit+
    ;

fragment
HexadecimalFractionalConstant
    :   HexadecimalDigitSequence? '.' HexadecimalDigitSequence
    |   HexadecimalDigitSequence '.'
    ;

fragment
BinaryExponentPart
    :   [pP] Sign? DigitSequence
    ;

fragment
HexadecimalDigitSequence
    :   HexadecimalDigit+
    ;

fragment
FloatingSuffix
    :   [flFL]
    ;

fragment
CharacterConstant
    :   '\'' CCharSequence '\''
    |   'L\'' CCharSequence '\''
    |   'u\'' CCharSequence '\''
    |   'U\'' CCharSequence '\''
    ;


fragment
CCharSequence
    :   CChar+
    ;

fragment
CChar
    :   ~['\\\r\n]
    |   EscapeSequence
    ;

fragment
EscapeSequence
    :   SimpleEscapeSequence
    |   OctalEscapeSequence
    |   HexadecimalEscapeSequence
    |   UniversalCharacterName
    ;

fragment
SimpleEscapeSequence
    :   '\\' ['"?abfnrtv\\]
    ;

fragment
OctalEscapeSequence
    :   '\\' OctalDigit OctalDigit? OctalDigit?
    ;

fragment
HexadecimalEscapeSequence
    :   '\\x' HexadecimalDigit+
    ;

StringLiteral
    :   EncodingPrefix? '"' SCharSequence? '"'
    ;

fragment
EncodingPrefix
    :   'u8'
    |   'u'
    |   'U'
    |   'L'
    ;

fragment
SCharSequence
    :   SChar+
    ;

fragment
SChar
    :   ~["\\\r\n]
    |   EscapeSequence
    |   '\\\n'   // Added line
    |   '\\\r\n' // Added line
    ;

fragment
NewlineFrag
    : '\r' '\n'?
    | '\n'
    ;

BlockComment
    :   '/*' .*? '*/'
        -> channel(HIDDEN)
    ;

LineComment
    :   '//' ~[\r\n]*
        -> channel(HIDDEN)
    ;
