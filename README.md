Bu, ağınızdaki açık bağlantı noktalarını bulmanızı sağlayan bir terminal uygulamasıdır.

*****************************************************************************************

Tüm bağlantı noktalarını tarar
En bilinen bağlantı noktalarını tarar
Arama zaman aşımını ayarlar
Belirli bir hedef belirler



Kurulum
**************************************************************************
Öncelikle Go ve Git'i indirmeniz gerekiyor
Not: Programın bazı özellikleri Git'i kullanıyor. Yüklediğinizden emin olun.
Daha sonra bu depoyu klonlayın:

git clone https://github.com/Tugce41/scanner-port.git
cd port-scanner



Termux'a yükleyin
***************************************************************
“`
$ pkg update -y && pkg upgrade -y   
“`
“`
$ pkg install git  
“`
“`
$ pkg install golang  
 “`
 “`
$ pkg install make  
“
“`
$ git clone https://github.com/Tugce41/scanner-port.git
“`
“`

$ cd port-scanner  
“`
$ make run  
“`




Herhangi bir yeri komut olarak kullanın
Örneğin Linux'ta bir terminal açın ve şu komutu yazın:

nano ~/.profile

Şimdi aşağıya doğru kaydırın ve şunu yapıştırın:

export PATH="$PATH:/usr/local/port-scanner/bin"

