# DNS-NOERROR

https://github.com/snowzach/dns-noerror

What is this? It's a simple DNS server that just returns no error for every request.
It's a hack... you shouldn't probably use it unless you know what you are doing.

Essentially, with a Mikrotik router you can run this container image and create a
static FWD DNS entry for your locally hosted domain and forward it to this container
and it will make your Mikrotik the authoritative server for this domain.
(You do need to make sure this is the last rule in your static entries)

## DNS Management Script

With that said, here's a script that automatically creates the FWD entries for the list of domains
and, when scheduled, will continue to ensure it's the last entry in your static entries.

```mikrotik
# Configuration - Edit these variables as needed
:local domains {"test.com";"whatever.net"}
:local comment "noerror"
:local forwardTo "192.168.1.10"

# Variables
:local domainsCount [:len $domains]

# Function to find the index position of an entry by its ID
:local findEntryIndex do={
    :for i from=0 to=($allEntriesCount - 1) do={
        :local checkEntryId [:pick $allEntries $i]
        :if ($checkEntryId = $entryId) do={
            :return $i
        }
    }
    :return -1
}

:foreach domain in=$domains do={
    # Update the local list of all DNS static entries
    :local allEntries [/ip/dns/static/find]
    :local allEntriesCount [:len $allEntries]

    # Check if entry exists
    :local searchComment ($comment . " - " . $domain)
    :local existingEntries [/ip/dns/static/find where comment=$searchComment]
    :if ([:len $existingEntries] = 0) do={
        :log info ("Forward entry for " . $domain . " does not exist, adding new entry.")
        /ip dns static add name=$domain match-subdomain=yes type=FWD forward-to=$forwardTo comment="$(comment) - $(domain)"
    } else={
        # Entry exists, get its actual position using our function
        :local entryId [:pick $existingEntries 0]
        :local entryPosition [$findEntryIndex allEntries=$allEntries allEntriesCount=$allEntriesCount entryId=$entryId]

        :if ($entryPosition < ($allEntriesCount - $domainsCount)) do={
            :log info ("Moving forward entry for domain " . $domain . " entry to end of list")
            /ip dns static move $entryId
        }
    }
}
```
