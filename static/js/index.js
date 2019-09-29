var editor = null;
$(document).ready(function() {

    $(".nav-link").click(function (e) {
        window.location = $( this ).attr('data-page');
    });

    var isMobile = window.matchMedia("only screen and (max-width: 760px)").matches;

    if (!isMobile) {
        var codeDiv = document.getElementById("wrap")
        var codeTextArea = document.getElementById('code');
        var codeText = codeTextArea.value;
        codeTextArea.outerHTML = '';
        codeDiv.id = "code";

        monaco.editor.defineTheme('coryPlayground', {
            base: 'vs-dark',
            inherit: true,
            rules: [],
            colors: {
                'editor.background': '#222222',
            }
        });

        editor = monaco.editor.create(document.getElementById('code'), {
            value: codeText,
            language: 'go',
            lineNumbers: "on",
            roundedSelection: true,
            scrollBeyondLastLine: false,
            readOnly: false,
            theme: "coryPlayground",
            minimap: {
                enabled: false
            }
        });
        // editor.layout();
    }

    playground({
        'codeEl':       '#code',
        'outputEl':     '#output',
        'runEl':        '#run, #embedRun',
        'fmtEl':        '#fmt',
        'fmtImportEl':  '#imports',
        'enableHistory': true,
        'enableShortcuts': true,
        'enableVet': true
    });
    playgroundEmbed({
        'codeEl':       '#code',
        'shareEl':      '#share',
        'embedEl':      '#embed',
        'embedLabelEl': '#embedLabel',
        'embedHTMLEl':  '#shareURL'
    });

    if (isMobile) {
        $('#code').linedtextarea();
        // Avoid line wrapping.
        $('#code').attr('wrap', 'off');
    }

    var about = $('#about');
    about.click(function(e) {
        if ($(e.target).is('a')) {
            return;
        }
        about.hide();
    });
    $('#aboutButton').click(function() {
        if (about.is(':visible')) {
            about.hide();
            return;
        }
        about.show();
    })
    // Preserve "Imports" checkbox value between sessions.
    if (readCookie('playgroundImports') == 'true') {
        $('#imports').attr('checked','checked');
    }
    $('#imports').change(function() {
        createCookie('playgroundImports', $(this).is(':checked') ? 'true' : '');
    });
});

function createCookie(name, value) {
    document.cookie = name+"="+value+"; path=/";
}

function readCookie(name) {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for(var i=0;i < ca.length;i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1,c.length);
        if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
    }
    return null;
}
