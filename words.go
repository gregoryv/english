package english

import "strings"

func Prepositions() []string {
	return strings.Fields(SinglePrepositionWords)
}

func Adverbs() []string {
	all := make([]string, 0)
	for _, words := range []string{
		HowAdverbWords,
		WhenAdverbWords,
		WhereAdverbWords,
		WhatExtentAdverbWords,
	} {
		all = append(all, strings.Fields(words)...)
	}
	return all
}

const (
	SinglePrepositionWords = `about beside near to above between of
towards across beyond off under after by on underneath against despite
onto unlike along down opposite until among during out up around
except outside upon as for over via at from past with before in round
within behind inside since without below into than beneath like
through`

	HowAdverbWords = `absentmindedly adoringly awkwardly beatifully
briskly brutally carefully cheerfully competitively eagerly
effortlessly extravagantly girlishly gracefully grimly happily
halfheartedly hungrily lazily lifelessly loyally quickly quitely
quizzically really recklessly remorsefully ruthlessly savagely
sloppily so stylishly unabashedly unevenly urgently well wishfully
worriedly`

	WhenAdverbWords = `after afterwards annually before daily never
now soon still then today tomorrow weekly when yesterday`

	WhereAdverbWords = `abroad anywhere away down everywhere here
home in inside out outside somewhere there underground upstairs`

	WhatExtentAdverbWords = `extremely not quite rather really
terribly too very`

	VerbWords = `accept accuse achieve acknowledge acquire adapt add
adjust admire admit adopt adore advise afford agree aim allow announce
anticipate apologize appear apply appreciate approach approve argue
arise arrange arrive ask assume assure astonish attach attempt attend
attract avoid awake bake bathe be bear beat become beg begin behave
believe belong bend bet bind bite blow boil borrow bounce bow break
breed bring broadcast build burn burst buy calculate can could care
carry catch celebrate change choose chop claim climb cling come commit
communicate compare compete complain complete concern confirm consent
consider consist consult contain continue convince cook cost count
crawl create creep criticize cry cut dance dare deal decide defer
delay deliver demand deny depend describe deserve desire destroy
determine develop differ disagree discover discuss dislike distribute
dive do doubt drag dream drill drink drive drop dry earn eat emphasize
enable encourage engage enhance enjoy ensure entail enter establish
examine exist expand expect experiment explain explore extend fail
fall feed feel fight find finish fit fly fold follow forbid forget
forgive freeze fry generate get give go grind grow hang happen hate
have hear hesitate hide hit hold hop hope hug hurry hurt identify
ignore illustrate imagine imply impress improve include incorporate
indicate inform insist install intend introduce invest investigate
involve iron jog jump justify keep kick kiss kneel knit know lack
laugh lay lead lean leap learn leave lend lie lift light lie like
listen look lose love maintain make manage matter may mean measure
meet melt mention might mind miss mix mow must need neglect negotiate
observe obtain occur offer open operate order organize ought overcome
overtake owe own paint participate pay peel perform persuade pinch
plan play point possess postpone pour practice prefer prepare pretend
prevent proceed promise propose protect prove pull punch pursue push
put qualify quit react read realize recall receive recollect recommend
reduce refer reflect refuse regret relate relax relieve rely remain
remember remind repair replace represent require resent resist retain
retire rid ride ring rise risk roast run sanction satisfy say scrub
see seem sell send serve set settle sew shake shall shed shine shoot
should show shrink shut sing sink sit ski sleep slice slide slip smell
snore solve sow speak specify spell spend spill spit spread squat
stack stand start steal stick sting stink stir stop stretch strike
struggle study submit succeed suffer suggest supply suppose surprise
survive swear sweep swell swim swing take talk taste teach tear tell
tend think threaten throw tiptoe tolerate translate try understand
vacuum value vary volunteer wait wake walk want warn wash watch wave
wear weep weigh whip will win wish would write`

	NounWords = `account act adjustment advertisement agreement air
amount amusement angle animal answer ant apparatus apple approval arch
argument arm army art attack attempt attention attraction authority
baby back bag balance ball band base basin basket bath bed bee
behavior belief bell berry bird birth bit bite blade blood blow board
boat body bone book boot bottle box boy brain brake branch brass bread
breath brick bridge brother brush bucket building bulb burn burst
business butter button cake camera canvas card care carriage cart cat
cause chain chalk chance change cheese chess chin church circle clock
cloth cloud coal coat collar color comb comfort committee company
comparison competition condition connection control cook copper copy
copy cord cork cough country cover cow crack credit crime crush cry
cup current curtain curve cushion damage danger daughter day death
debt decision degree design desire destruction detail development
digestion direction discovery discussion disease disgust distance
distribution division dog door doubt drain drawer dress drink driving
drop dust ear earth edge education effect egg end engine error event
example exchange existence expansion experience expert eye face fact
fall family farm father fear feather feeling fiction field fight
finger fire fish flag flame flight floor flower fly fold food foot
force fork form fowl frame friend front fruit garden girl glass glove
goat gold government grain grass grip group growth guide gun hair
hammer hand harbor harmony hat hate head hearing heart heat help
history hole hook hope horn horse hospital hour house humor ice idea
impulse increase industry ink insect instrument insurance interest
invention iron island jelly jewel join journey judge jump kettle key
kick kiss knee knife knot knowledge land language laugh lead leaf
learning leather leg letter level library lift light limit line linen
lip liquid list lock look loss love low machine man manager map mark
market mass match meal measure meat meeting memory metal middle milk
mind mine minute mist money monkey month moon morning mother motion
mountain mouth move muscle music nail name nation neck need needle
nerve net news night noise nose note number nut observation offer
office oil operation opinion orange order organization ornament oven
owner page pain paint paper parcel part paste payment peace pen pencil
person picture pig pin pipe place plane plant plate play pleasure
plough pocket point poison polish porter position pot potato powder
power price print prison process produce profit property prose protest
pull pump punishment purpose push quality question rail rain range rat
rate ray reaction reading reason receipt record regret relation
religion representative request respect rest reward rhythm rice ring
river road rod roll roof room root rub rule run sail salt sand scale
school science scissors screw sea seat secretary seed selection self
sense servant sex shade shake shame sheep shelf ship shirt shock shoe
side sign silk silver sister size skin skirt sky sleep slip slope
smash smell smile smoke snake sneeze snow soap society sock son song
sort sound soup space spade sponge spoon spring square stage stamp
star start statement station steam steel stem step stick stitch
stocking stomach stone stop store story street stretch structure
substance sugar suggestion summer sun support surprise swim system
table tail talk taste tax teaching tendency test theory thing thought
thread throat thumb thunder ticket time tin toe tongue tooth top touch
town trade train transport tray tree trick trouble trousers turn twist
umbrella unit use value verse vessel view voice walk wall war wash
waste watch water wave wax way weather week weight wheel whip whistle
wind window wine wing winter wire woman wood wool word work worm wound
writing year`

	AdjectiveWords = `adorable adventurous aggressive agreeable
alert alive amused angry annoyed annoying anxious arrogant ashamed
attractive average awful bad beautiful better bewildered black bloody
blue blue-eyed blushing bored brainy brave breakable bright busy calm
careful cautious charming cheerful clean clear clever cloudy clumsy
colorful combative comfortable concerned condemned confused
cooperative courageous crazy creepy crowded cruel curious cute
dangerous dark dead defeated defiant delightful depressed determined
different difficult disgusted distinct disturbed dizzy doubtful drab
dull eager easy elated elegant embarrassed enchanting encouraging
energetic enthusiastic envious evil excited expensive exuberant fair
faithful famous fancy fantastic fierce filthy fine foolish fragile
frail frantic friendly frightened funny gentle gifted glamorous
gleaming glorious good gorgeous graceful grieving grotesque grumpy
handsome happy healthy helpful helpless hilarious homeless homely
horrible hungry hurt ill important impossible inexpensive innocent
inquisitive itchy jealous jittery jolly joyous kind light lively
lonely long lovely lucky magnificent misty modern motionless muddy
mushy mysterious nasty naughty nervous nice nutty obedient obnoxious
odd old-fashioned open outrageous outstanding panicky perfect plain
pleasant poised poor powerful precious prickly proud putrid puzzled
quaint real relieved repulsive rich scary selfish shiny shy silly
sleepy smiling smoggy sore sparkling splendid spotless stormy strange
stupid successful super talented tame tasty tender tense terrible
thankful thoughtful thoughtless tired tough troubled ugliest ugly
uninterested unsightly unusual upset uptight vast victorious vivacious
wandering weary wicked wide-eyed wild witty worried worrisome wrong
zany zealous`
)
