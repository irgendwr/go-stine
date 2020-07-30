package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const ascii = `
                                                      .,,*,,,,,,,.                                               
                                                 .,****,'.     .',,,,.                                           
                                              .,**,                  ,,,.                                        
                                            ,**'    ',,,,'..           '*,.                                      
                                          ,*,.  ,****,                   .*,                                     
                                        .*,  .*T*,                         ,,                                    
                                        T,  *i,                             ,*                                   
                       .,              ** .*.                                ,*                                  
                     .,,i* .',,.      'S.                                     T*                                 
                   ,,,,iT,',',i.      is                    ,S,               *S                                 
                  ,' ',' ',*eNi       S* .'.  .',,,,,,''. .*SS,               *N'                                
               ,,,*    ',,.*,.*       E, ,NEeSStiItsTiiii*,' ,,               *N,                                
             ,*,.    ,T,  ,' '*      .E. ,'.           .,,,'',Ii***,          iN,                                
            **       ,T,..,  i'      ,i  *'',,,*,      ,*,,,,,,''.,T,         sN,                                
            i          *T*' .s       *,, *,,**Ti*'   ..  ,*NNn,.   ,,         nN,                                
            i    .en,  .ENI,*I      .'i, ,*,,*NE,.    ,.  .is,.    ,*         tE'                                
           'S.   .Ei.  .ENNNNi        E, '* '','      ,,           ,*         TE.                                
           .n.    ,  ,*ENNNNS'       ,E.  *           .,           ,T         *E.                                
           .I.    .EENNNNNt,         ni   *,                        *.    *i  ,E'                                
           ,*     ,NNNNe*.          *E.   .*       ','',,           *,    *I*' e,                                
           *      *NNS,            ,E,     .*.                       ,    ** ,*sT                                
          ,,      SNS.            ,S,       '*.      ...'..,         ,*   *,   *S'                               
         .,      'EN,            ,T. .,,,,   ,T,   ',,,,',*,        ,i**. '*    ..                               
         ,.      *Ni            ,t****,,,..,Is'ii'  ,,,,,,        .*ei .*, *                                     
        ,,       IN,                  *T****'   ,T*'            ,**,,S.  ',**                                    
       .*       'NN                   .           ,**,.    .,,,*,''''i*    '**                                   
       *'       *Nn                               .  .,,,,,*,'''''''',T      .                                   
      *,        SN*                            '**             .'''''',si*,                                      
     ,i        ,NE.                         ',iSIT*.              .'',*IneNET*,                                  
    .t'        iNi                      .,*seENNEeIsi*,,'      ..',*isISENNNNNNESi*,.                            
    T*        .EN,                   '*TeEENNNNNNNEEeSIssTiiiiiTsssISeENNNNNNNNNNNNNNes*,.                       
   ,t         *NS               .,*iSEENNNNNNNNNNNNNNNNEEEeeeeeeEEENNNNNNNNNNNNNNNNNNNNNNEt*.                    
  .n,        .EN*           .,*seENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNEi'                  
  **         sNe.   .,,*****,*eENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNT'                
 .n.        ,NNe******,.      *SENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNT'              
 **        .ENI,              'TSENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNEiENNNNNNNNNNNNNNNNNNNNNNNNNi.            
 i.        iN*                 *seNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNnIEi.*EeSNNNNNNNNNNNNNNNNNNNNNNNE,           
'*         ,.                  'TsENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNe,,n* ,e,,ENNNNNNNNNNNNNNNNNNNNNNNN*          
,.                              *sSNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNn  s* ,e'.eNNNNNNNNNNNNNNNNNNNNNNNEI.         
*'                          ',*IEntENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNET,,,,,,,,,sNNNNNNNNNNNNNNNNNNNNNEi'.,         
*e.                 .,**seENNNNNNSsENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN*   ,*,   iNNNNNNNNNNNNNNNNNNe*,    .*.       
.Ee.         ',**sSENNNNNNNNEeIT**ENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNS.   INE.  'ENNNNNNNNNNNNNNEI,        .*'      
 ,eEi**iitSEENNNNNNEeIT**,,.      iNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNeTTTTENEITTTeNNNNNNNNNNNEs*.           .*,     
   ,*iiii****,,,,'.               .SNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNES*,.                **    
                                   ,ENNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNeieE*.                   *i.  
                                    \STiNEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE/   \ei,                  ,s, 
`

// asciiCmd represents the ascii command
var asciiCmd = &cobra.Command{
	Use:   "ascii",
	Short: "STiNE ascii art",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(ascii)
	},
}

func init() {
	rootCmd.AddCommand(asciiCmd)
}
