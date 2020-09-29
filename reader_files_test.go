// Generated by reader_util_test.go.
// Only edit to bootstrap new entries or change existing entries.

package fit_test

var decodeTestFiles = [...]struct {
	folder      string
	name        string
	wantErr     bool
	fingerprint uint64
	compress    bool
	dopts       testingDecodeOpts
	skipEncode  bool
	encodeNote  string
}{
	{
		"me",
		"activity-small-fenix2-run.fit",
		false,
		4159419868942230381,
		true,
		tdoAllWithDiscardLogger,
		true,
		"Decode mismatch due to unknown fields",
	},
	{
		"fitsdk",
		"Activity.fit",
		false,
		17200987073157828903,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"MonitoringFile.fit",
		false,
		17438554766263628503,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"fitsdk",
		"Settings.fit",
		false,
		13777716833885766673,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WeightScaleMultiUser.fit",
		false,
		5002902390457186100,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutCustomTargetValues.fit",
		false,
		1795447305759253807,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutIndividualSteps.fit",
		false,
		13432157117047365566,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutRepeatGreaterThanStep.fit",
		false,
		10036597401811670605,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WorkoutRepeatSteps.fit",
		false,
		2000919626000542509,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"fitsdk",
		"WeightScaleSingleUser.fit",
		false,
		170466418772016754,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"garmin-edge-500-activitiy.fit",
		false,
		6593572577964555599,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"sample-activity-indoor-trainer.fit",
		false,
		16971736118512546826,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"python-fitparse",
		"compressed-speed-distance.fit",
		false,
		0,
		false,
		tdoNone,
		false,
		"",
	},
	{
		"python-fitparse",
		"antfs-dump.63.fit",
		false,
		17636932664957499573,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"sram",
		"Settings.fit",
		false,
		15334527618026463003,
		true,
		tdoNone,
		true,
		"Fails due to encoder using profile length for arrays (#37)",
	},
	{
		"sram",
		"Settings2.fit",
		false,
		3203160144706376883,
		true,
		tdoNone,
		true,
		"Fails due to encoder using profile length for arrays (#37)",
	},
	{
		"dcrainmaker",
		"Edge810-Vector-2013-08-16-15-35-10.fit",
		false,
		7445399672663916413,
		true,
		tdoNone,
		true,
		"Fails because first message has different valid fields (#36)",
	},
	{
		"misc",
		"2013-02-06-12-11-14.fit",
		false,
		13418237519260555178,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"misc",
		"2015-10-13-08-43-15.fit",
		false,
		9824244170689694008,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"corrupt",
		"activity-filecrc.fit",
		true,
		1128179132232601357,
		true,
		tdoNone,
		false,
		"",
	},
	{
		"corrupt",
		"activity-unexpected-eof.fit",
		true,
		3268441259118371812,
		true,
		tdoNone,
		false,
		"",
	},
}
