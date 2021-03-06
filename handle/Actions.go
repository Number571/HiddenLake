package handle

import (
	"github.com/number571/gopeer"
	"github.com/number571/hiddenlake/settings"
)

func Actions(client *gopeer.Client, pack *gopeer.Package) {
	client.HandleAction(settings.TITLE_EMAIL, pack, getEmail, setEmail)
	client.HandleAction(settings.TITLE_ARCHIVE, pack, getArchive, setArchive)
	client.HandleAction(settings.TITLE_TESTCONN, pack, getTestconn, setTestconn)
	client.HandleAction(settings.TITLE_GROUPCHAT, pack, getGroupchat, setGroupchat)
	client.HandleAction(settings.TITLE_PRIVATECHAT, pack, getPrivatechat, setPrivatechat)
}
