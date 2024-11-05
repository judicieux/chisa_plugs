package l9plugins

import (
        "github.com/LeakIX/l9format"
        "github.com/judicieux/chisa_plugs/web"
)

var webPlugins = []l9format.WebPluginInterface{
		web.AwsCredentialsHttpPlugin{},
		web.BootstrapHttpPlugin{},
		web.CloudConfigHttpPlugin{},
		web.ConfigBakHttpPlugin{},
		web.ConfigEnvRubyHttpPlugin{},
		web.ConfigIncHttpPlugin{},
		web.ConfigJsonHttp{},
		web.CodeJsonReply{},
		web.ConfigPhpDistHttpPlugin{},
		web.ConfigRubyHttpPlugin{},
		web.ConfigStorageHttpPlugin{},
		web.ConfigYmlHttpPlugin{},
		web.DashboardPhpInfoHttpPlugin{},
		web.DockerComposeProdHttpPlugin{},
		web.DotEnvHttpPlugin{},
		web.EnvBackupHttpPlugin{},
		web.EnvBakHttpPlugin{},
		web.EnvDevHttpPlugin{},
		web.EnvExampleHttpPlugin{},
		web.EnvLiveHttpPlugin{},
		web.EnvLocalHttpPlugin{},
		web.EnvOldHttpPlugin{},
		web.EnvProdHttpPlugin{},
		web.EnvProductionHttpPlugin{},
		web.EnvRcHttpPlugin{},
		web.EnvSampleHttpPlugin{},
		web.EnvSaveHttpPlugin{},
		web.EnvStageHttpPlugin{},
		web.DotGitlabciHttpPlugin{},
		web.DotNetrcPluginHttp{},
		web.GitConfigHttpPlugin{},
		web.ConfigAssetsHttpPlugin{},
		web.ConfigContentHttpPlugin{},
		web.DataConfigHttpPlugin{},
		web.ConfigImagesHttpPlugin{},
		web.ConfigJsHttpPlugin{},
		web.ConfigLibHttpPlugin{},
		web.ConfigMediaHttpPlugin{},
		web.PublicConfigHttpPlugin{},
		web.ServerConfigHttpPlugin{},
		web.SrcConfigHttpPlugin{},
		web.ConfigStaticHttpPlugin{},
		web.ConfigSubPathApiHttpPlugin{},
		web.ConfigSubPathDevHttpPlugin{},
		web.ConfigSubPathJsHttpPlugin{},
		web.ConfigSubPathMediaHttpPlugin{},
		web.ConfigSubPathWwwHttpPlugin{},
		web.LinusadminPhpInfoHttpPlugin{},
		web.OldPhpInfoHttpPlugin{},
		web.PhpInfoHttpPlugin{},
		web.PInfoHttpPlugin{},
		web.ProfilerPhpInfoHttpPlugin{},
		web.TiretPhpInfoHttpPlugin{},
		web.TravisYmlHttpPlugin{},
		web.UnderscorePhpInfoHttpPlugin{},
}

func GetWebPlugins() []l9format.WebPluginInterface {
        return webPlugins
}
