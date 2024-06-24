package DocTrim

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var testXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:wpc="http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
    xmlns:o="urn:schemas-microsoft-com:office:office"
    xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"
    xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math"
    xmlns:v="urn:schemas-microsoft-com:vml"
    xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
    xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
    xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup"
    xmlns:wpi="http://schemas.microsoft.com/office/word/2010/wordprocessingInk"
    xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main"
    xmlns:wne="http://schemas.microsoft.com/office/word/2006/wordml"
    xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape"
    xmlns:w10="urn:schemas-microsoft-com:office:word"
    xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing"
    xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"
    xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml"
    xmlns:w16cex="http://schemas.microsoft.com/office/word/2018/wordml/cex"
    xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid"
    xmlns:w16="http://schemas.microsoft.com/office/word/2018/wordml"
    xmlns:w16sdtdh="http://schemas.microsoft.com/office/word/2020/wordml/sdtdatahash"
    xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex"
    mc:Ignorable="w14 w15 w16se w16cid w16 w16cex w16sdtdh wp14">
    <w:body>
        <w:p>
            <w:pPr>
                <w:keepNext w:val="false" />
                <w:keepLines w:val="false" />
                <w:pageBreakBefore w:val="false" />
                <w:pBdr></w:pBdr>
                <w:snapToGrid w:val="true" />
                <w:spacing w:line="240" w:lineRule="auto" />
                <w:ind />
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:pPr>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia" w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                    <w:lang w:val="en-US" w:eastAsia="zh-CN" />
                </w:rPr>
                <w:t xml:space="preserve">4</w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">．用配方法解方程</w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
            <m:oMath>
                <m:sSup>
                    <m:sSupPr>
                        <m:ctrlPr></m:ctrlPr>
                    </m:sSupPr>
                    <m:e>
                        <m:r>
                            <w:rPr>
                                <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                    w:eastAsia="Cambria Math" />
                                <w:sz w:val="24" />
                            </w:rPr>
                            <m:rPr>
                                <m:sty m:val="i" />
                            </m:rPr>
                            <m:t>x</m:t>
                        </m:r>
                    </m:e>
                    <m:sup>
                        <m:r>
                            <w:rPr>
                                <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                    w:eastAsia="Cambria Math" />
                                <w:sz w:val="14" />
                            </w:rPr>
                            <m:rPr>
                                <m:sty m:val="p" />
                            </m:rPr>
                            <m:t>2</m:t>
                        </m:r>
                    </m:sup>
                </m:sSup>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>+</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>4</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="i" />
                    </m:rPr>
                    <m:t>x</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>+</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>1</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>=</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>0</m:t>
                </m:r>
            </m:oMath>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">时，配方后所得的方程是</w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia" w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">（     ）</w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
        </w:p>
        <w:p>
            <w:pPr>
                <w:keepNext w:val="false" />
                <w:keepLines w:val="false" />
                <w:pageBreakBefore w:val="false" />
                <w:pBdr></w:pBdr>
                <w:snapToGrid w:val="true" />
                <w:spacing w:line="240" w:lineRule="auto" />
                <w:ind w:firstLine="420" />
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:pPr>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia" w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">A．</w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
            <m:oMath>
                <m:sSup>
                    <m:sSupPr>
                        <m:ctrlPr></m:ctrlPr>
                    </m:sSupPr>
                    <m:e>
                        <m:d>
                            <m:dPr>
                                <m:begChr m:val="(" />
                                <m:endChr m:val=")" />
                                <m:ctrlPr></m:ctrlPr>
                            </m:dPr>
                            <m:e>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="i" />
                                    </m:rPr>
                                    <m:t>x</m:t>
                                </m:r>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="p" />
                                    </m:rPr>
                                    <m:t>+</m:t>
                                </m:r>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="p" />
                                    </m:rPr>
                                    <m:t>2</m:t>
                                </m:r>
                            </m:e>
                        </m:d>
                    </m:e>
                    <m:sup>
                        <m:r>
                            <w:rPr>
                                <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                    w:eastAsia="Cambria Math" />
                                <w:sz w:val="14" />
                            </w:rPr>
                            <m:rPr>
                                <m:sty m:val="p" />
                            </m:rPr>
                            <m:t>2</m:t>
                        </m:r>
                    </m:sup>
                </m:sSup>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>=</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>3</m:t>
                </m:r>
            </m:oMath>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">     </w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia" w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">B．</w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
            <m:oMath>
                <m:sSup>
                    <m:sSupPr>
                        <m:ctrlPr></m:ctrlPr>
                    </m:sSupPr>
                    <m:e>
                        <m:d>
                            <m:dPr>
                                <m:begChr m:val="(" />
                                <m:endChr m:val=")" />
                                <m:ctrlPr></m:ctrlPr>
                            </m:dPr>
                            <m:e>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="i" />
                                    </m:rPr>
                                    <m:t>x</m:t>
                                </m:r>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="p" />
                                    </m:rPr>
                                    <m:t>−</m:t>
                                </m:r>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="p" />
                                    </m:rPr>
                                    <m:t>2</m:t>
                                </m:r>
                            </m:e>
                        </m:d>
                    </m:e>
                    <m:sup>
                        <m:r>
                            <w:rPr>
                                <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                    w:eastAsia="Cambria Math" />
                                <w:sz w:val="14" />
                            </w:rPr>
                            <m:rPr>
                                <m:sty m:val="p" />
                            </m:rPr>
                            <m:t>2</m:t>
                        </m:r>
                    </m:sup>
                </m:sSup>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>=</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>3</m:t>
                </m:r>
            </m:oMath>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">    </w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia" w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">C．</w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
            <m:oMath>
                <m:sSup>
                    <m:sSupPr>
                        <m:ctrlPr></m:ctrlPr>
                    </m:sSupPr>
                    <m:e>
                        <m:d>
                            <m:dPr>
                                <m:begChr m:val="(" />
                                <m:endChr m:val=")" />
                                <m:ctrlPr></m:ctrlPr>
                            </m:dPr>
                            <m:e>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="i" />
                                    </m:rPr>
                                    <m:t>x</m:t>
                                </m:r>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="p" />
                                    </m:rPr>
                                    <m:t>+</m:t>
                                </m:r>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="p" />
                                    </m:rPr>
                                    <m:t>2</m:t>
                                </m:r>
                            </m:e>
                        </m:d>
                    </m:e>
                    <m:sup>
                        <m:r>
                            <w:rPr>
                                <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                    w:eastAsia="Cambria Math" />
                                <w:sz w:val="14" />
                            </w:rPr>
                            <m:rPr>
                                <m:sty m:val="p" />
                            </m:rPr>
                            <m:t>2</m:t>
                        </m:r>
                    </m:sup>
                </m:sSup>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>=</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>5</m:t>
                </m:r>
            </m:oMath>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">      </w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia" w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
                <w:t xml:space="preserve">D．</w:t>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
            <m:oMath>
                <m:sSup>
                    <m:sSupPr>
                        <m:ctrlPr></m:ctrlPr>
                    </m:sSupPr>
                    <m:e>
                        <m:d>
                            <m:dPr>
                                <m:begChr m:val="(" />
                                <m:endChr m:val=")" />
                                <m:ctrlPr></m:ctrlPr>
                            </m:dPr>
                            <m:e>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="i" />
                                    </m:rPr>
                                    <m:t>x</m:t>
                                </m:r>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="p" />
                                    </m:rPr>
                                    <m:t>+</m:t>
                                </m:r>
                                <m:r>
                                    <w:rPr>
                                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                            w:eastAsia="Cambria Math" />
                                        <w:sz w:val="24" />
                                    </w:rPr>
                                    <m:rPr>
                                        <m:sty m:val="p" />
                                    </m:rPr>
                                    <m:t>4</m:t>
                                </m:r>
                            </m:e>
                        </m:d>
                    </m:e>
                    <m:sup>
                        <m:r>
                            <w:rPr>
                                <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                                    w:eastAsia="Cambria Math" />
                                <w:sz w:val="14" />
                            </w:rPr>
                            <m:rPr>
                                <m:sty m:val="p" />
                            </m:rPr>
                            <m:t>2</m:t>
                        </m:r>
                    </m:sup>
                </m:sSup>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>=</m:t>
                </m:r>
                <m:r>
                    <w:rPr>
                        <w:rFonts w:ascii="Cambria Math" w:hAnsi="Cambria Math"
                            w:eastAsia="Cambria Math" />
                        <w:sz w:val="24" />
                    </w:rPr>
                    <m:rPr>
                        <m:sty m:val="p" />
                    </m:rPr>
                    <m:t>5</m:t>
                </m:r>
            </m:oMath>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
            <w:r>
                <w:rPr>
                    <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
                    <w:szCs w:val="21" />
                </w:rPr>
            </w:r>
        </w:p>
        <w:p>
            <w:pPr>
                <w:pBdr></w:pBdr>
                <w:spacing />
                <w:ind />
                <w:rPr></w:rPr>
            </w:pPr>
            <w:r></w:r>
            <w:r></w:r>
        </w:p>
        <w:p>
            <w:pPr>
                <w:pBdr></w:pBdr>
                <w:spacing />
                <w:ind />
                <w:rPr></w:rPr>
            </w:pPr>
            <w:r></w:r>
            <w:r></w:r>
        </w:p>
        <w:sectPr>
            <w:footnotePr></w:footnotePr>
            <w:endnotePr></w:endnotePr>
            <w:type w:val="nextPage" />
            <w:pgSz w:h="16838" w:orient="portrait" w:w="11906" />
            <w:pgMar w:top="1134" w:right="850" w:bottom="1134" w:left="1701" w:header="709"
                w:footer="709" w:gutter="0" />
            <w:cols w:num="1" w:sep="0" w:space="708" w:equalWidth="1"></w:cols>
            <w:docGrid w:type="default" w:linePitch="360" w:charSpace="0" />
        </w:sectPr>
    </w:body>
</w:document>
`

func TestCompact(t *testing.T) {
	// 1. 创建一个Slim实例
	s := DocTrim{}

	// 2. 创建一个Node实例
	data, err := s.Pack(bytes.NewReader([]byte(testXml)))
	// 3. 调用Node的Compact方法
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%d -> %d", len(testXml), len(data))
	// 4. 验证Node的Compact方法是否正确
	to, err := s.Unpack(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if !EqualXml([]byte(testXml), to) {
		t.Fatal("Not equals")
	}
}

func OneTestFile(t *testing.T, filename string) {
	// 1. 创建一个Slim实例
	s := DocTrim{}

	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file, err := os.Open(filename)
	// Get file info

	if err != nil {
		t.Fatal(err)
		return
	}
	// 2. 创建一个Node实例
	data, err := s.Pack(file)
	// 3. 调用Node的Compact方法
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%s %d -> %d (%4.2f%%)", filename, fileInfo.Size(), len(data), float64(len(data))/float64(fileInfo.Size())*100.0)
	// 4. 验证Node的Compact方法是否正确
	copyData, err := s.Unpack(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%d -> %d", len(data), len(copyData))

	file.Seek(0, 0)
	fromData, _ := ioutil.ReadAll(file)
	if !EqualXml(fromData, copyData) {
		t.Fatal("Not equals")
	}
}

func TestAll(t *testing.T) {
	OneTestFile(t, "docs/document.xml")
}

func TestOne(t *testing.T) {
	OneTestFile(t, "docs/test.xml")
}

func TestText(t *testing.T) {
	OneTestFile(t, "docs/text.xml")
}

func TestEmptyToSelfClosing(t *testing.T) {
	input := `<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"><w:p><w:pPr><w:pBdr hash="1"></w:pBdr><w:tabs><w:tab w:val="left" w:leader="none" w:pos="2835"></w:tab></w:tabs><w:spacing hash="4"></w:spacing><w:ind w:firstLine="709"></w:ind><w:jc w:val="left"></w:jc><w:rPr hash="7"></w:rPr></w:pPr><w:r><w:tab hash="9"></w:tab><w:tab ref="9"></w:tab></w:r><w:r hash="b"></w:r></w:p></w:document>`
	s := DocTrim{}
	data, err := s.Pack(bytes.NewReader([]byte(input)))
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%s", string(data))
	log.Printf("%d -> %d", len(input), len(data))
}
