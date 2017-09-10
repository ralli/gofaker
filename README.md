

# gofaker
`import "github.com/ralli/gofaker"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package gofaker helps generating test data like persons names, addresses, email addresses, etc.

### Credits
gofaker is (more or less) a port of Benjamin Curtis (stympy) faker library to generate
fake data (which is itself a port of Perl's Data::Faker library).

### Locales
Faker supports generating fake data in different languages. You can find a list of all available locales in the example of the function AllLocales().

### Thread safety
Currently Faker objects are not thread safe since they use a single random number generator instance. Genrating
fake data from multiple goroutines concurrently is still possible when using a separate Faker instance per goroutine.




## <a name="pkg-index">Index</a>
* [func AllLocales() []string](#AllLocales)
* [func Romanize(s string) string](#Romanize)
* [func ToSnake(in string) string](#ToSnake)
* [type Address](#Address)
  * [func (a *Address) BuildingNumber() string](#Address.BuildingNumber)
  * [func (a *Address) City() string](#Address.City)
  * [func (a *Address) Country() string](#Address.Country)
  * [func (a *Address) CountryCode() string](#Address.CountryCode)
  * [func (a *Address) Latitude() string](#Address.Latitude)
  * [func (a *Address) Longitude() string](#Address.Longitude)
  * [func (a *Address) State() string](#Address.State)
  * [func (a *Address) StateAbbr() string](#Address.StateAbbr)
  * [func (a *Address) StreetAddress() string](#Address.StreetAddress)
  * [func (a *Address) StreetName() string](#Address.StreetName)
  * [func (a *Address) TimeZone() string](#Address.TimeZone)
  * [func (a *Address) ZipCode() string](#Address.ZipCode)
  * [func (a *Address) ZipCodeByState(stateName string) string](#Address.ZipCodeByState)
* [type Code](#Code)
  * [func (code *Code) Base10ISBN() string](#Code.Base10ISBN)
  * [func (code *Code) Base13ISBN() string](#Code.Base13ISBN)
* [type Commerce](#Commerce)
  * [func (commerce *Commerce) Color() string](#Commerce.Color)
  * [func (commerce *Commerce) Department() string](#Commerce.Department)
  * [func (commerce *Commerce) Price() float64](#Commerce.Price)
  * [func (commerce *Commerce) ProductName() string](#Commerce.ProductName)
* [type Company](#Company)
  * [func (company *Company) Bullshit() string](#Company.Bullshit)
  * [func (company *Company) DUNSNumber() string](#Company.DUNSNumber)
  * [func (company *Company) EIN() string](#Company.EIN)
  * [func (company *Company) Industry() string](#Company.Industry)
  * [func (company *Company) Name() string](#Company.Name)
  * [func (company *Company) SwedishOrganisationNumber() string](#Company.SwedishOrganisationNumber)
* [type Faker](#Faker)
  * [func NewFaker(locale string) (*Faker, error)](#NewFaker)
  * [func (faker *Faker) Bothify(s string) string](#Faker.Bothify)
  * [func (faker *Faker) Fetch(key string) (string, error)](#Faker.Fetch)
  * [func (faker *Faker) FetchList(key string) ([][]string, error)](#Faker.FetchList)
  * [func (faker *Faker) FetchOrEmpty(key string) string](#Faker.FetchOrEmpty)
  * [func (faker *Faker) Letterify(s string) string](#Faker.Letterify)
  * [func (faker *Faker) MustFetch(key string) string](#Faker.MustFetch)
  * [func (faker *Faker) MustFetchList(key string) [][]string](#Faker.MustFetchList)
  * [func (faker *Faker) MustParse(key string) string](#Faker.MustParse)
  * [func (faker *Faker) Numerify(s string) string](#Faker.Numerify)
  * [func (faker *Faker) Parse(key string) (string, error)](#Faker.Parse)
  * [func (faker *Faker) RandomDigit() rune](#Faker.RandomDigit)
  * [func (faker *Faker) RandomLetter() rune](#Faker.RandomLetter)
  * [func (faker *Faker) Reset()](#Faker.Reset)
  * [func (faker *Faker) Sample(arr []string, num int) []string](#Faker.Sample)
* [type Internet](#Internet)
  * [func (internet *Internet) DomainWord() string](#Internet.DomainWord)
  * [func (internet *Internet) Email() string](#Internet.Email)
  * [func (internet *Internet) EmailWithUsername(name string) string](#Internet.EmailWithUsername)
  * [func (internet *Internet) FreeMail() string](#Internet.FreeMail)
  * [func (internet *Internet) FreeMailWithName(name string) string](#Internet.FreeMailWithName)
  * [func (internet *Internet) IPv4Address() string](#Internet.IPv4Address)
  * [func (internet *Internet) IPv6Address() string](#Internet.IPv6Address)
  * [func (internet *Internet) MacAddress() string](#Internet.MacAddress)
  * [func (internet *Internet) Password() string](#Internet.Password)
  * [func (internet *Internet) PasswordWithLength(minLength int, maxLength int) string](#Internet.PasswordWithLength)
  * [func (internet *Internet) PrivateIPv4Address() string](#Internet.PrivateIPv4Address)
  * [func (internet *Internet) SafeMail() string](#Internet.SafeMail)
  * [func (internet *Internet) SafeMailWithName(name string) string](#Internet.SafeMailWithName)
  * [func (internet *Internet) UserName() string](#Internet.UserName)
  * [func (internet *Internet) UserWithName(specifier string) string](#Internet.UserWithName)
* [type Lorem](#Lorem)
  * [func (lorem *Lorem) Characters(num int) string](#Lorem.Characters)
  * [func (lorem *Lorem) Paragraph() string](#Lorem.Paragraph)
  * [func (lorem *Lorem) ParagraphNum(numSentences int, randomSentencesToAdd int) string](#Lorem.ParagraphNum)
  * [func (lorem *Lorem) Paragraphs(numParagraphs int) []string](#Lorem.Paragraphs)
  * [func (lorem *Lorem) Sentence() string](#Lorem.Sentence)
  * [func (lorem *Lorem) SentenceNum(numWords int, randomWordsToAdd int) string](#Lorem.SentenceNum)
  * [func (lorem *Lorem) SentencesNum(numSentences int) string](#Lorem.SentencesNum)
  * [func (lorem *Lorem) Word() string](#Lorem.Word)
  * [func (lorem *Lorem) Words() string](#Lorem.Words)
  * [func (lorem *Lorem) WordsNum(numWords int) string](#Lorem.WordsNum)
  * [func (lorem *Lorem) WordsSupplemental(numWords int) string](#Lorem.WordsSupplemental)
* [type Name](#Name)
  * [func (name *Name) FirstName() string](#Name.FirstName)
  * [func (name *Name) LastName() string](#Name.LastName)
  * [func (name *Name) Name() string](#Name.Name)
  * [func (name *Name) Title() string](#Name.Title)
* [type Phone](#Phone)
  * [func (phone *Phone) CellPhone() string](#Phone.CellPhone)
  * [func (phone *Phone) PhoneNumber() string](#Phone.PhoneNumber)
  * [func (phone *Phone) SubscriberNumber() string](#Phone.SubscriberNumber)

#### <a name="pkg-examples">Examples</a>
* [Package](#example_)
* [Address.BuildingNumber](#example_Address_BuildingNumber)
* [Address.City](#example_Address_City)
* [Address.Country](#example_Address_Country)
* [Address.CountryCode](#example_Address_CountryCode)
* [Address.Latitude](#example_Address_Latitude)
* [Address.Longitude](#example_Address_Longitude)
* [Address.State](#example_Address_State)
* [Address.StateAbbr](#example_Address_StateAbbr)
* [Address.StreetAddress](#example_Address_StreetAddress)
* [Address.StreetName](#example_Address_StreetName)
* [Address.TimeZone](#example_Address_TimeZone)
* [Address.ZipCode](#example_Address_ZipCode)
* [Address.ZipCodeByState](#example_Address_ZipCodeByState)
* [AllLocales](#example_AllLocales)
* [Commerce.Color](#example_Commerce_Color)
* [Commerce.Department](#example_Commerce_Department)
* [Commerce.Price](#example_Commerce_Price)
* [Commerce.ProductName](#example_Commerce_ProductName)
* [Company.Bullshit](#example_Company_Bullshit)
* [Company.DUNSNumber](#example_Company_DUNSNumber)
* [Company.EIN](#example_Company_EIN)
* [Company.Industry](#example_Company_Industry)
* [Company.Name](#example_Company_Name)
* [Company.SwedishOrganisationNumber](#example_Company_SwedishOrganisationNumber)
* [Lorem.Characters](#example_Lorem_Characters)
* [Lorem.Paragraphs](#example_Lorem_Paragraphs)
* [Lorem.Sentence](#example_Lorem_Sentence)
* [Lorem.SentencesNum](#example_Lorem_SentencesNum)
* [Lorem.Word](#example_Lorem_Word)
* [Lorem.Words](#example_Lorem_Words)
* [Lorem.WordsNum](#example_Lorem_WordsNum)
* [Lorem.WordsSupplemental](#example_Lorem_WordsSupplemental)
* [Name.FirstName](#example_Name_FirstName)
* [Name.LastName](#example_Name_LastName)
* [Name.Name](#example_Name_Name)
* [Name.Title](#example_Name_Title)
* [Phone.PhoneNumber](#example_Phone_PhoneNumber)
* [Phone.SubscriberNumber](#example_Phone_SubscriberNumber)
* [Romanize](#example_Romanize)

#### <a name="pkg-files">Package files</a>
[address.go](/src/github.com/ralli/gofaker/address.go) [code.go](/src/github.com/ralli/gofaker/code.go) [commerce.go](/src/github.com/ralli/gofaker/commerce.go) [company.go](/src/github.com/ralli/gofaker/company.go) [doc.go](/src/github.com/ralli/gofaker/doc.go) [faker.go](/src/github.com/ralli/gofaker/faker.go) [fakerdata.go](/src/github.com/ralli/gofaker/fakerdata.go) [internet.go](/src/github.com/ralli/gofaker/internet.go) [lorem.go](/src/github.com/ralli/gofaker/lorem.go) [name.go](/src/github.com/ralli/gofaker/name.go) [phone.go](/src/github.com/ralli/gofaker/phone.go) [stringutils.go](/src/github.com/ralli/gofaker/stringutils.go) 





## <a name="AllLocales">func</a> [AllLocales](/src/target/faker.go?s=6447:6473#L257)
``` go
func AllLocales() []string
```
AllLocales returns a list of all available locales supported by gofaker.



## <a name="Romanize">func</a> [Romanize](/src/target/stringutils.go?s=1217:1247#L16)
``` go
func Romanize(s string) string
```
Romanize maps characters from a given input string to the ascii range. Currently supports
cyrillic characters and german umlauts.



## <a name="ToSnake">func</a> [ToSnake](/src/target/faker.go?s=992:1022#L40)
``` go
func ToSnake(in string) string
```
ToSnake converts the given string to snake case following the Golang format:
acronyms are converted to lower-case and preceded by an underscore.




## <a name="Address">type</a> [Address](/src/target/address.go?s=107:144#L1)
``` go
type Address struct {
    // contains filtered or unexported fields
}
```
Address provides functions to generate postal address related fake data.










### <a name="Address.BuildingNumber">func</a> (\*Address) [BuildingNumber](/src/target/address.go?s=664:705#L16)
``` go
func (a *Address) BuildingNumber() string
```
BuildingNumber generates a building number.




### <a name="Address.City">func</a> (\*Address) [City](/src/target/address.go?s=179:210#L1)
``` go
func (a *Address) City() string
```
City generates a cities name.




### <a name="Address.Country">func</a> (\*Address) [Country](/src/target/address.go?s=1628:1662#L46)
``` go
func (a *Address) Country() string
```
Country generates a countries name like "Bulgaria".




### <a name="Address.CountryCode">func</a> (\*Address) [CountryCode](/src/target/address.go?s=1781:1819#L51)
``` go
func (a *Address) CountryCode() string
```
Generates a two letter ISO Country Code like "DE" for "Germany".




### <a name="Address.Latitude">func</a> (\*Address) [Latitude](/src/target/address.go?s=1938:1973#L56)
``` go
func (a *Address) Latitude() string
```
Latitude generates the latitude of a geographical position.




### <a name="Address.Longitude">func</a> (\*Address) [Longitude](/src/target/address.go?s=2107:2143#L61)
``` go
func (a *Address) Longitude() string
```
Longitude generates the longitude of a geographical position.




### <a name="Address.State">func</a> (\*Address) [State](/src/target/address.go?s=1492:1524#L41)
``` go
func (a *Address) State() string
```
State generates a states name like "Arizona".




### <a name="Address.StateAbbr">func</a> (\*Address) [StateAbbr](/src/target/address.go?s=1353:1389#L36)
``` go
func (a *Address) StateAbbr() string
```
StateAbbr generates a abbreviated state name like "AZ" for "Arizona".




### <a name="Address.StreetAddress">func</a> (\*Address) [StreetAddress](/src/target/address.go?s=501:541#L11)
``` go
func (a *Address) StreetAddress() string
```
StreetAddress generates a street name with locale dependent additional data (a building number for example).




### <a name="Address.StreetName">func</a> (\*Address) [StreetName](/src/target/address.go?s=297:334#L6)
``` go
func (a *Address) StreetName() string
```
StreetName generates a street name.




### <a name="Address.TimeZone">func</a> (\*Address) [TimeZone](/src/target/address.go?s=1192:1227#L31)
``` go
func (a *Address) TimeZone() string
```
TimeZone generates a time zone name like "Europe/Berlin".




### <a name="Address.ZipCode">func</a> (\*Address) [ZipCode](/src/target/address.go?s=814:848#L21)
``` go
func (a *Address) ZipCode() string
```
ZipCode generates a zip code.




### <a name="Address.ZipCodeByState">func</a> (\*Address) [ZipCodeByState](/src/target/address.go?s=984:1041#L26)
``` go
func (a *Address) ZipCodeByState(stateName string) string
```
ZipCodeByState generates a zipcode with an attached state name.




## <a name="Code">type</a> [Code](/src/target/code.go?s=105:139#L1)
``` go
type Code struct {
    // contains filtered or unexported fields
}
```
Code provides functions to generate codes (ISBN numbers and the like).










### <a name="Code.Base10ISBN">func</a> (\*Code) [Base10ISBN](/src/target/code.go?s=329:366#L2)
``` go
func (code *Code) Base10ISBN() string
```
Base10ISBN generates an old (10 digit) ISBN number (international standard book number).
Users normally would like to generate a Base13ISBN. Since 2007 all ISBNs have to be base 13.




### <a name="Code.Base13ISBN">func</a> (\*Code) [Base13ISBN](/src/target/code.go?s=719:756#L22)
``` go
func (code *Code) Base13ISBN() string
```
Base10ISBN generates an ISBN number (international standard book number).




## <a name="Commerce">type</a> [Commerce](/src/target/commerce.go?s=118:156#L1)
``` go
type Commerce struct {
    // contains filtered or unexported fields
}
```
Commerce provides functions to generate fake data for products sold in a web shop.










### <a name="Commerce.Color">func</a> (\*Commerce) [Color](/src/target/commerce.go?s=221:261#L1)
``` go
func (commerce *Commerce) Color() string
```
Color generates the name of a products color like "yellow".




### <a name="Commerce.Department">func</a> (\*Commerce) [Department](/src/target/commerce.go?s=409:454#L6)
``` go
func (commerce *Commerce) Department() string
```
Department generates a category or department name like "Books", "Movies" or "Electronics".




### <a name="Commerce.Price">func</a> (\*Commerce) [Price](/src/target/commerce.go?s=877:918#L16)
``` go
func (commerce *Commerce) Price() float64
```
Price generates a price in the range from 0.00 to 99.99.




### <a name="Commerce.ProductName">func</a> (\*Commerce) [ProductName](/src/target/commerce.go?s=564:610#L11)
``` go
func (commerce *Commerce) ProductName() string
```
ProductName generates a random product name.




## <a name="Company">type</a> [Company](/src/target/company.go?s=132:169#L1)
``` go
type Company struct {
    // contains filtered or unexported fields
}
```
Company provides functions to generate fake data related to companies.










### <a name="Company.Bullshit">func</a> (\*Company) [Bullshit](/src/target/company.go?s=605:646#L16)
``` go
func (company *Company) Bullshit() string
```
Bullshit generates a bullshit motto like "implement value-added web-readyness".
See <a href="http://dack.com/web/bullshit.html">http://dack.com/web/bullshit.html</a> for details.




### <a name="Company.DUNSNumber">func</a> (\*Company) [DUNSNumber](/src/target/company.go?s=1014:1057#L31)
``` go
func (company *Company) DUNSNumber() string
```
DUNSNumber generates a DUNS number.




### <a name="Company.EIN">func</a> (\*Company) [EIN](/src/target/company.go?s=888:924#L26)
``` go
func (company *Company) EIN() string
```
EIN Generates an Employers ID number (EIN).




### <a name="Company.Industry">func</a> (\*Company) [Industry](/src/target/company.go?s=369:410#L10)
``` go
func (company *Company) Industry() string
```
Industry generates the companies industry like "Space and Defense".




### <a name="Company.Name">func</a> (\*Company) [Name](/src/target/company.go?s=207:244#L5)
``` go
func (company *Company) Name() string
```
Name generates a companies name.




### <a name="Company.SwedishOrganisationNumber">func</a> (\*Company) [SwedishOrganisationNumber](/src/target/company.go?s=1179:1237#L36)
``` go
func (company *Company) SwedishOrganisationNumber() string
```
SwedishOrganisationNumber generates a swedish organisation number.




## <a name="Faker">type</a> [Faker](/src/target/faker.go?s=113:319#L3)
``` go
type Faker struct {
    Name     *Name
    Address  *Address
    Phone    *Phone
    Internet *Internet
    Company  *Company
    Code     *Code
    Commerce *Commerce
    Lorem    *Lorem
    // contains filtered or unexported fields
}
```
Faker does something







### <a name="NewFaker">func</a> [NewFaker](/src/target/faker.go?s=3914:3958#L160)
``` go
func NewFaker(locale string) (*Faker, error)
```
NewFaker creates a new Faker instance for a given locale





### <a name="Faker.Bothify">func</a> (\*Faker) [Bothify](/src/target/faker.go?s=5503:5547#L222)
``` go
func (faker *Faker) Bothify(s string) string
```
Bothify replaces all hash signs within a string with random digits and all




### <a name="Faker.Fetch">func</a> (\*Faker) [Fetch](/src/target/faker.go?s=2767:2820#L115)
``` go
func (faker *Faker) Fetch(key string) (string, error)
```
Fetch gets a fake data value without further expansion. That means, the result may
contain unexpanded variables like '#{last_name}'




### <a name="Faker.FetchList">func</a> (\*Faker) [FetchList](/src/target/faker.go?s=3541:3602#L143)
``` go
func (faker *Faker) FetchList(key string) ([][]string, error)
```



### <a name="Faker.FetchOrEmpty">func</a> (\*Faker) [FetchOrEmpty](/src/target/faker.go?s=3414:3465#L135)
``` go
func (faker *Faker) FetchOrEmpty(key string) string
```
FetchOrEmpty gets a fake data value without further expansion. That means, the result may
contain unexpanded variables like '#{last_name}'. Returns an empty string on error.




### <a name="Faker.Letterify">func</a> (\*Faker) [Letterify](/src/target/faker.go?s=5213:5259#L209)
``` go
func (faker *Faker) Letterify(s string) string
```
Letterify replaces all question marks within a string with random letters.




### <a name="Faker.MustFetch">func</a> (\*Faker) [MustFetch](/src/target/faker.go?s=3109:3157#L125)
``` go
func (faker *Faker) MustFetch(key string) string
```
MustFetch gets a fake data value without further expansion. That means, the result may
contain unexpanded variables like '#{last_name}'. Panics on error.




### <a name="Faker.MustFetchList">func</a> (\*Faker) [MustFetchList](/src/target/faker.go?s=3717:3773#L151)
``` go
func (faker *Faker) MustFetchList(key string) [][]string
```



### <a name="Faker.MustParse">func</a> (\*Faker) [MustParse](/src/target/faker.go?s=2504:2552#L105)
``` go
func (faker *Faker) MustParse(key string) string
```
MustParse returns the value of Parse and panics in case of an error




### <a name="Faker.Numerify">func</a> (\*Faker) [Numerify](/src/target/faker.go?s=4780:4825#L191)
``` go
func (faker *Faker) Numerify(s string) string
```
Numerify replaces all hash signs within a string with random digits.




### <a name="Faker.Parse">func</a> (\*Faker) [Parse](/src/target/faker.go?s=2109:2162#L88)
``` go
func (faker *Faker) Parse(key string) (string, error)
```
Parse gets fake data given a key and expands variable values recursively.




### <a name="Faker.RandomDigit">func</a> (\*Faker) [RandomDigit](/src/target/faker.go?s=4626:4664#L186)
``` go
func (faker *Faker) RandomDigit() rune
```
RandomDigit returns a random digit.




### <a name="Faker.RandomLetter">func</a> (\*Faker) [RandomLetter](/src/target/faker.go?s=5041:5080#L204)
``` go
func (faker *Faker) RandomLetter() rune
```
RandomLetter returns a random letter (a-z, A-Z)




### <a name="Faker.Reset">func</a> (\*Faker) [Reset](/src/target/faker.go?s=4531:4558#L181)
``` go
func (faker *Faker) Reset()
```
Reset resets the fakers random number generator so that a repeatable sequence of values is returned.
Useful for testing.




### <a name="Faker.Sample">func</a> (\*Faker) [Sample](/src/target/faker.go?s=6166:6224#L246)
``` go
func (faker *Faker) Sample(arr []string, num int) []string
```
Sample returns a sample (some random items) of a given array. The sample entries are not guaranteed to be unique.

faker.Sample([]string{"a", "b", "c", "d", "e", "f"}, 6) might return []string{"e", "a", "e", "b", "f", "c"}




## <a name="Internet">type</a> [Internet](/src/target/internet.go?s=552:590#L13)
``` go
type Internet struct {
    // contains filtered or unexported fields
}
```
Internet provides functions to generate fake data for email addresses, user names, URLs and the like.










### <a name="Internet.DomainWord">func</a> (\*Internet) [DomainWord](/src/target/internet.go?s=3687:3732#L88)
``` go
func (internet *Internet) DomainWord() string
```
DomainWord returns a valid domain part of an URL.




### <a name="Internet.Email">func</a> (\*Internet) [Email](/src/target/internet.go?s=2200:2240#L60)
``` go
func (internet *Internet) Email() string
```
Email returns an email address




### <a name="Internet.EmailWithUsername">func</a> (\*Internet) [EmailWithUsername](/src/target/internet.go?s=1972:2035#L55)
``` go
func (internet *Internet) EmailWithUsername(name string) string
```
EmailWithUserName returns an email-Address for a given name




### <a name="Internet.FreeMail">func</a> (\*Internet) [FreeMail](/src/target/internet.go?s=2900:2943#L72)
``` go
func (internet *Internet) FreeMail() string
```
FreeMail generates an email address for a random user at a known free mail provider like gmail.
The given user name will be normalized so that it becomes a valid email user name.




### <a name="Internet.FreeMailWithName">func</a> (\*Internet) [FreeMailWithName](/src/target/internet.go?s=2555:2617#L66)
``` go
func (internet *Internet) FreeMailWithName(name string) string
```
FreeMailWithName generates an email address for a given user at a known free mail provider like gmail.
The given user name will be normalized so that it becomes a valid email user name.




### <a name="Internet.IPv4Address">func</a> (\*Internet) [IPv4Address](/src/target/internet.go?s=4578:4624#L128)
``` go
func (internet *Internet) IPv4Address() string
```
IPv4Address generates an IP-Address.




### <a name="Internet.IPv6Address">func</a> (\*Internet) [IPv6Address](/src/target/internet.go?s=5206:5252#L156)
``` go
func (internet *Internet) IPv6Address() string
```
IPv6Address returns a valid IP V6 Address.




### <a name="Internet.MacAddress">func</a> (\*Internet) [MacAddress](/src/target/internet.go?s=4327:4372#L119)
``` go
func (internet *Internet) MacAddress() string
```
MacAddress generates a MAC address.




### <a name="Internet.Password">func</a> (\*Internet) [Password](/src/target/internet.go?s=1817:1860#L50)
``` go
func (internet *Internet) Password() string
```
Password returns a password




### <a name="Internet.PasswordWithLength">func</a> (\*Internet) [PasswordWithLength](/src/target/internet.go?s=1434:1515#L39)
``` go
func (internet *Internet) PasswordWithLength(minLength int, maxLength int) string
```
PasswordWithLength returns a password with a length ranging from minLength to maxLength




### <a name="Internet.PrivateIPv4Address">func</a> (\*Internet) [PrivateIPv4Address](/src/target/internet.go?s=4887:4940#L137)
``` go
func (internet *Internet) PrivateIPv4Address() string
```
PrivateIPv4Address generates an IP Address reserved for the local network like "192.168.x.y".




### <a name="Internet.SafeMail">func</a> (\*Internet) [SafeMail](/src/target/internet.go?s=3495:3538#L83)
``` go
func (internet *Internet) SafeMail() string
```
SafeMailWithName generates an email for which it is safe to send test mail to (like example.com).




### <a name="Internet.SafeMailWithName">func</a> (\*Internet) [SafeMailWithName](/src/target/internet.go?s=3228:3290#L78)
``` go
func (internet *Internet) SafeMailWithName(name string) string
```
SafeMailWithName generates an email address for which it is safe to send test mail to (like example.com).
The given user name will be normalized so that it becomes a valid email user name.




### <a name="Internet.UserName">func</a> (\*Internet) [UserName](/src/target/internet.go?s=980:1023#L26)
``` go
func (internet *Internet) UserName() string
```
UserName returns a random username.




### <a name="Internet.UserWithName">func</a> (\*Internet) [UserWithName](/src/target/internet.go?s=687:750#L18)
``` go
func (internet *Internet) UserWithName(specifier string) string
```
UserWithName transforms a users name into something that can be used as a username (login).




## <a name="Lorem">type</a> [Lorem](/src/target/lorem.go?s=113:148#L1)
``` go
type Lorem struct {
    // contains filtered or unexported fields
}
```
Lorem provides functions for fake data related to "Lorem Ipsum" sentences.










### <a name="Lorem.Characters">func</a> (\*Lorem) [Characters](/src/target/lorem.go?s=1083:1129#L24)
``` go
func (lorem *Lorem) Characters(num int) string
```
Characters genrates a sequence of alphanumeric characters.




### <a name="Lorem.Paragraph">func</a> (\*Lorem) [Paragraph](/src/target/lorem.go?s=2762:2800#L76)
``` go
func (lorem *Lorem) Paragraph() string
```
Paragraph returns a paragraph.




### <a name="Lorem.ParagraphNum">func</a> (\*Lorem) [ParagraphNum](/src/target/lorem.go?s=2550:2633#L71)
``` go
func (lorem *Lorem) ParagraphNum(numSentences int, randomSentencesToAdd int) string
```
ParagraphNum returns a paragraph with a minimum number of sentences. There may be a random number of sentences
which will be added to the paragraph.




### <a name="Lorem.Paragraphs">func</a> (\*Lorem) [Paragraphs](/src/target/lorem.go?s=2883:2941#L81)
``` go
func (lorem *Lorem) Paragraphs(numParagraphs int) []string
```
Paragraphs returns a list of paragraphs.




### <a name="Lorem.Sentence">func</a> (\*Lorem) [Sentence](/src/target/lorem.go?s=2046:2083#L56)
``` go
func (lorem *Lorem) Sentence() string
```
Returns a sentence. A sentence starts with an uppercase first
word and ends with a period.




### <a name="Lorem.SentenceNum">func</a> (\*Lorem) [SentenceNum](/src/target/lorem.go?s=1582:1656#L38)
``` go
func (lorem *Lorem) SentenceNum(numWords int, randomWordsToAdd int) string
```
SentenceNum generates a sentence with a minimum number of words. A sentence starts with an uppercase first
word and ends with a period. The number of words may vary by a random number of additional words.




### <a name="Lorem.SentencesNum">func</a> (\*Lorem) [SentencesNum](/src/target/lorem.go?s=2174:2231#L61)
``` go
func (lorem *Lorem) SentencesNum(numSentences int) string
```
SentencesNum returns a given number of sentences.




### <a name="Lorem.Word">func</a> (\*Lorem) [Word](/src/target/lorem.go?s=195:228#L1)
``` go
func (lorem *Lorem) Word() string
```
Word generates a single Lorem Ipsum word.




### <a name="Lorem.Words">func</a> (\*Lorem) [Words](/src/target/lorem.go?s=341:375#L6)
``` go
func (lorem *Lorem) Words() string
```
Words generates three Lorem words separated by whitespace.




### <a name="Lorem.WordsNum">func</a> (\*Lorem) [WordsNum](/src/target/lorem.go?s=477:526#L11)
``` go
func (lorem *Lorem) WordsNum(numWords int) string
```
Words generates a given number of lorem words separated by spaces.




### <a name="Lorem.WordsSupplemental">func</a> (\*Lorem) [WordsSupplemental](/src/target/lorem.go?s=762:820#L17)
``` go
func (lorem *Lorem) WordsSupplemental(numWords int) string
```
WordsSupplemental generates a given number of lorem words using an extendes set of words separated by spaces.




## <a name="Name">type</a> [Name](/src/target/name.go?s=71:105#L1)
``` go
type Name struct {
    // contains filtered or unexported fields
}
```
Name generates fake data related to persons names.










### <a name="Name.FirstName">func</a> (\*Name) [FirstName](/src/target/name.go?s=144:180#L1)
``` go
func (name *Name) FirstName() string
```
FirstName generates a first name.




### <a name="Name.LastName">func</a> (\*Name) [LastName](/src/target/name.go?s=269:304#L4)
``` go
func (name *Name) LastName() string
```
LastName generates a last name.




### <a name="Name.Name">func</a> (\*Name) [Name](/src/target/name.go?s=394:425#L9)
``` go
func (name *Name) Name() string
```
Name returns a full persons name.




### <a name="Name.Title">func</a> (\*Name) [Title](/src/target/name.go?s=528:560#L14)
``` go
func (name *Name) Title() string
```
Title returns a job title like "Senior Vice Chief".




## <a name="Phone">type</a> [Phone](/src/target/phone.go?s=72:107#L1)
``` go
type Phone struct {
    // contains filtered or unexported fields
}
```
Phone provides functions to generate phone numbers.










### <a name="Phone.CellPhone">func</a> (\*Phone) [CellPhone](/src/target/phone.go?s=418:456#L4)
``` go
func (phone *Phone) CellPhone() string
```
CellPhone genrates valid cell phone numbers.




### <a name="Phone.PhoneNumber">func</a> (\*Phone) [PhoneNumber](/src/target/phone.go?s=248:288#L1)
``` go
func (phone *Phone) PhoneNumber() string
```
PhoneNumber generates a phone number. Different phone numbers will be generated in different formats. The formats are locale dependent.




### <a name="Phone.SubscriberNumber">func</a> (\*Phone) [SubscriberNumber](/src/target/phone.go?s=586:631#L9)
``` go
func (phone *Phone) SubscriberNumber() string
```
SubscriberNumber generates subscriber numbers.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
