let unless =  macro(condition, consequence, alternative) {
     quote(
        if (!(unquote(condition))) { 
                unquote(consequence); 
        }
        else { 
            unquote(alternative); 
        }
    ); 
};

unless(10>5, puts("not greater"), puts("greater"));
unless(10>50, puts("not greater"), puts("greater"));