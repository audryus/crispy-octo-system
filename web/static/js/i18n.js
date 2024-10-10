i18next.use(i18nextHttpBackend).init({
    lng: 'en',
    load: 'languageOnly',
    debug: true,
    fallbackLng: 'en',
    backend: {
        loadPath: '/static/locales/{{lng}}/translation.json'
    }
  }, function(err, t) {
    if (err) {
        throw err
    }
    localize = locI18next.init(i18next);
    localize('[data-i18n]');
  });