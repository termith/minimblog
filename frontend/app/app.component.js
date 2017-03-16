(function(app) {
  app.NewPostComponent =
    ng.core.Component({
      selector: 'new-post',
      templateUrl: 'templates/new_post_template.html'
    })
    .Class({
      constructor: function() {}
    });
})(window.app || (window.app = {}));
