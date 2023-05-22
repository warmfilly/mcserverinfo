# Maintainer: Your Name <youremail@domain.com>
pkgname=mcserverinfo
pkgver=1
pkgrel=1
pkgdesc="d"
url=https://github.com/warmfilly/mcserverinfo
license=('unknown')
source_x86_64=("mcserverinfo-${pkgver}-linux-amd64.tar.gz::${url}/archive/refs/tags/v${pkgver}.tar.gz")
arch=('x86_64')
provides=('mcserverinfo')
sha256sums_x86_64=('da3f97212e0849681653adbe25fd6efbbf1c417d08d5952d29df3f510fb89d18')

build() {
    cd "$pkgname-$pkgver"
    go build
}

package() {
    pwd
    echo "$pkgdir"
    echo "Path: $pkgdir/usr/bin/mcserverinfo"
    cd "$pkgname-$pkgver"
    pwd
    mkdir -p "$pkgdir/usr/bin/mcserver"
    install mcserverinfo "$pkgdir/usr/bin/mcserverinfo"
}
