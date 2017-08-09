(function(app) {
  app.AppModule =
    ng.core.NgModule({
      imports: [ ng.platformBrowser.BrowserModule ],
      declarations: [ app.NewPostComponent ],
      bootstrap: [ app.NewPostComponent ]
    })
    .Class({
      constructor: function() {}
    });
})(window.app || (window.app = {}));
