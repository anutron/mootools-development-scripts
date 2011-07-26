#!/bin/bash
sync() {
    cd lib/configurations;
    for f in *
    do
      cd $f
      echo "~~~~~~~~~~~~ Syncronizing project: $f ~~~~~~~~~~~~"
      ../../../env/bin/crepo sync;
      cd ..
    done
}


check() {
    cd lib/configurations;
    for f in *
    do
      cd $f
      echo "~~~~~~~~~~~~ Checking project: $f ~~~~~~~~~~~~"
      ../../../env/bin/crepo check-dirty;
      cd ..
    done
}
apps(){
    cd lib/applications
    ../../env/bin/crepo check-dirty
}

if [ $# -lt 1 ]; then
  echo "Choose an option: (i)nstall, (d)ependency check, (r)un server on port 9876, (l)ive - run the server on port 80, (s)ync libraries, or (c) check for repos for unstaged changes."
  read option
else
  option=$1
fi
case $option in
    [i]* ) 
        cd lib/applications/mootools-runner;
        git submodule update --init;
        cd ../dev-app;
        ln -s ../../../settings.py;
        cd ../depender/django;
        ../../../../env/bin/python setup.py develop;
        cd ../../dev-app;
        ../../../env/bin/python setup.py develop;;
    [d]* )
        cd lib/applications/dev-app;
        ../../../env/bin/python manage.py depender_check;;
    [r]* )
        cd lib/applications/dev-app;
        ../../../env/bin/python manage.py runserver_plus 0.0.0.0:9876;;
    [l]* )
        cd lib/applications/dev-app;
        ../../../env/bin/python manage.py runserver 0.0.0.0:80;;
    [c]* )
        check;;
    [a]* )
        apps;;
    [s]* )
        sync;;
    * ) echo "Please choose from i, d, r, l, or s.";;
esac


